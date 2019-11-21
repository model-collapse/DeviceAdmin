package components

import (
	"fmt"
	"html/template"
	"math"
	"we/device_admin/models/usage"

	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	dc "github.com/GoAdminGroup/themes/adminlte/components/description"
	pg "github.com/GoAdminGroup/themes/adminlte/components/progress_group"
)

type UsageChart struct {
	usage *usage.ClusterUsage
}

func NewUsageChart(usage *usage.ClusterUsage) *UsageChart {
	return &UsageChart{usage: usage}
}

func (c *UsageChart) getLabels(history []usage.SnapShot) (ret []string) {
	ret = make([]string, len(history))
	for i, sn := range history {
		ret[i] = sn.Loc.Format("15:04:05")
	}

	return
}

func (c *UsageChart) getCPUCurve(history []usage.SnapShot) (ret []float64) {
	ret = make([]float64, len(history))
	for i, sn := range history {
		ret[i] = float64(sn.CPU)
	}

	return
}

func (c *UsageChart) getGPUCurve(history []usage.SnapShot) (ret []float64) {
	ret = make([]float64, len(history))
	for i, sn := range history {
		ret[i] = float64(sn.GPU)
	}

	return
}

func (c *UsageChart) getMemCurve(history []usage.SnapShot) (ret []float64) {
	ret = make([]float64, len(history))
	for i, sn := range history {
		ret[i] = float64(sn.Mem)
	}

	return
}

func (c *UsageChart) getDeviceNumCurve(history []usage.SnapShot) (ret []float64) {
	ret = make([]float64, len(history))
	for i, sn := range history {
		ret[i] = float64(sn.DeviceNum)
	}

	return
}

func (c *UsageChart) GetContent() template.HTML {
	components := template2.Get(config.Get().Theme)
	colComp := components.Col()
	history := c.usage.GetStatusHistory()
	chart := chartjs.Line().
		SetID("usage_chart").
		SetHeight(180).
		SetTitle("Cluster Usage Summary").
		SetLabels(c.getLabels(history)).
		AddDataSet("CPU").
		DSData(c.getCPUCurve(history)).
		DSFill(false).
		DSCubicInterpolationMode("monotone").
		DSBorderColor("rgba(102, 176, 255, 0.9)").
		DSLineTension(0.1).DSPointStyle("line").
		AddDataSet("GPU").
		DSData(c.getGPUCurve(history)).
		DSFill(false).
		DSCubicInterpolationMode("monotone").
		DSBorderColor("rgba(255, 176, 102, 0.9)").
		DSLineTension(0.1).DSPointStyle("line").
		AddDataSet("Memory").
		DSData(c.getMemCurve(history)).
		DSFill(false).
		DSCubicInterpolationMode("monotone").
		DSBorderColor("rgba(176, 255, 102, 0.9)").
		DSLineTension(0.1).DSPointStyle("line").
		GetContent()

	current := history[len(history)-1]
	progCPU := pg.New().SetTitle("CPU Usage").SetColor("#76b2d4").SetDenominator(100).SetMolecular(int(100 * current.CPU)).SetPercent(int(100 * current.CPU)).GetContent()
	progGPU := pg.New().SetTitle("GPU Usage").SetColor("#ace0ae").SetDenominator(100).SetMolecular(int(100 * current.GPU)).SetPercent(int(100 * current.GPU)).GetContent()
	progMem := pg.New().SetTitle("Memory Usage").SetColor("#fdd698").SetDenominator(100).SetMolecular(int(100 * current.Mem)).SetPercent(int(100 * current.Mem)).GetContent()

	ccol := colComp.SetSize(map[string]string{"md": "8"}).SetContent(chart).GetContent()
	pcol := colComp.SetContent(template.HTML(`<p class="text-center"><strong>Current Usage</strong></p>`) + progCPU + progGPU + progMem).
		SetSize(map[string]string{"md": "4"}).
		GetContent()

	crow := components.Row().SetContent(ccol + pcol).GetContent()

	var last usage.SnapShot
	if len(history) > 1 {
		last = history[len(history)-1]
	}

	snapDiff := usage.SnapshotDiff(last, current)
	arrow := func(v int) string {
		if v > 0 {
			return "up"
		}

		return "down"
	}

	arrowColor := func(v int) template.HTML {
		if v > 0 {
			return "green"
		}

		return "red"
	}

	percentMD := func(a, b int) template.HTML {
		v := math.Abs(float64(a) / float64(b))
		vv := int(v * 100)
		return template.HTML(fmt.Sprintf("%d", vv))
	}

	dCPUCores := dc.New().SetTitle("Total CPU Cores").SetNumber(template.HTML(fmt.Sprintf("%d", current.CPUCores))).SetBorder("right").SetArrow(arrow(snapDiff.CPUCores)).SetPercent(percentMD(snapDiff.CPUCores, current.CPUCores)).SetColor(arrowColor(snapDiff.CPUCores))
	dGPUCores := dc.New().SetTitle("Total GPU Cores").SetNumber(template.HTML(fmt.Sprintf("%d", current.GPUCores))).SetBorder("right").SetArrow(arrow(snapDiff.GPUCores)).SetPercent(percentMD(snapDiff.GPUCores, current.GPUCores)).SetColor(arrowColor(snapDiff.GPUCores))
	dMemCap := dc.New().SetTitle("Total Memory Capacity (MB)").SetNumber(template.HTML(fmt.Sprintf("%d MB", current.MemoryCap))).SetBorder("right").SetArrow(arrow(snapDiff.MemoryCap)).SetPercent(percentMD(snapDiff.MemoryCap, current.MemoryCap)).SetColor(arrowColor(snapDiff.MemoryCap))
	dEdgeDevice := dc.New().SetTitle("#Edge Device").SetNumber(template.HTML(fmt.Sprintf("%d", current.EdgeNum))).SetBorder("right").SetArrow(arrow(snapDiff.EdgeNum)).SetPercent(percentMD(snapDiff.EdgeNum, current.EdgeNum)).SetColor(arrowColor(snapDiff.EdgeNum))
	dSensorDevice := dc.New().SetTitle("#Sensor Device").SetNumber(template.HTML(fmt.Sprintf("%d", current.SensorNum))).SetBorder("right").SetArrow(arrow(snapDiff.SensorNum)).SetPercent(percentMD(snapDiff.SensorNum, current.SensorNum)).SetColor(arrowColor(snapDiff.SensorNum))

	size2 := map[string]string{"sm": "3", "xs": "6"}
	dCPUCoresCol := colComp.SetContent(dCPUCores.GetContent()).SetSize(size2).GetContent()
	dGPUCoresCol := colComp.SetContent(dGPUCores.GetContent()).SetSize(size2).GetContent()
	dMemCapCol := colComp.SetContent(dMemCap.GetContent()).SetSize(size2).GetContent()
	dEdgeDeviceCol := colComp.SetContent(dEdgeDevice.GetContent()).SetSize(size2).GetContent()
	dSensorDeviceCol := colComp.SetContent(dSensorDevice.GetContent()).SetSize(size2).GetContent()

	drow := components.Row().SetContent(dCPUCoresCol + dGPUCoresCol + dMemCapCol + dEdgeDeviceCol + dSensorDeviceCol).GetContent()

	box := components.Box().WithHeadBorder(true).SetHeader("Cluster Usage Summary").SetBody(crow).SetFooter(drow).GetContent()
	col := colComp.SetSize(map[string]string{"md": "12"}).SetContent(box).GetContent()

	return components.Row().SetContent(col).GetContent()
}
