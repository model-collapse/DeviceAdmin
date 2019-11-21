package components

import (
	"fmt"
	"html/template"
	"log"
	"sort"
	"we/device_admin/models/usage"

	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"

	//"github.com/GoAdminGroup/themes/adminlte/components/smallbox"
	"we/device_admin/dashboard/themes/devicebox"
)

type DeviceUnits struct {
	usage *usage.ClusterUsage
}

func NewDeviceUnits(usage *usage.ClusterUsage) *DeviceUnits {
	return &DeviceUnits{usage: usage}
}

func toPercent(f float32) int {
	return int(f * 100)
}

func (d *DeviceUnits) makeStatusTitle(stat usage.DeviceStatus) (ret template.HTML) {
	var rets string
	if stat.HealthyScore() < 0 {
		rets = "Inactive"
	} else {
		switch stat.DeviceTypeID {
		case usage.RaspberryPI_3B, usage.RaspberryPI_4B, usage.RaspberryPI_Zero:
			rets = fmt.Sprintf("<font size=\"3\">CPU: </font>%d<sup style=\"font-size: 12px\">%%</sup>&nbsp;&nbsp;<font size=\"3\">Mem: </font>%d<sup style=\"font-size: 12px\">%%</sup>", toPercent(stat.CPU), toPercent(stat.Mem))
		case usage.JetsonNano:
			tmpl := `<table>
						<tr><td style="font-size:12px;padding-right: 2px;padding-bottom: 2px" valign="bottom">CPU:</td><td style="font-size:21px" valign="center">%d<sup style="font-size:12px">%%</sup></td><td rowspan="2" style="font-size:18px;padding-right: 6px;padding-bottom: 6px;padding-left: 10px" valign="bottom">Mem:</td><td style="font-size:38px" valign="center" rowspan="2">%d<sup style="font-size:18px">%%</sup></td></tr>
						<tr><td style="font-size:12px;padding-right: 2px;padding-bottom: 2px" valign="bottom">GPU:</td><td style="font-size:21px" valign="center">%d<sup style="font-size:12px">%%</sup></td></tr>
			         </table>
			`
			rets = fmt.Sprintf(tmpl, toPercent(stat.CPU), toPercent(stat.Mem), toPercent(stat.GPU))
		case usage.SparkfunBLE:
			rets = fmt.Sprintf("Active")
		case usage.ESP32CAM:
			rets = fmt.Sprintf("Active")
		}
	}

	return template.HTML(rets)
}

func healthyScoreToColor(s int) string {
	if s > 3 {
		log.Printf("no such healthy score %d", s)
		return "white"
	} else if s < 0 {
		return "gray"
	}

	ar := []string{"blue", "red", "orange", "green"}
	return ar[s]
}

func (d *DeviceUnits) GetContent() template.HTML {
	devices := d.usage.GetDeiveStatusMap()

	components := template2.Get(config.Get().Theme)
	colComp := components.Col()
	var size = map[string]string{"md": "3", "sm": "6"}
	var colsEdge template.HTML
	var colsSensor template.HTML

	names := make([]string, 0, len(devices))
	for name := range devices {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		stat := devices[name]
		sbox := devicebox.New().
			SetValue(d.makeStatusTitle(stat)).
			SetColor(template.HTML(healthyScoreToColor(stat.HealthyScore()))).
			SetTitle(template.HTML(fmt.Sprintf("<strong>%s</strong>&nbsp;&nbsp;|&nbsp;&nbsp;<font size=\"3\">ID: %s</font>", stat.DeviceTypeStr, name))).
			SetIcon(IconMap[stat.DeviceTypeID]).
			GetContent()

		cnt := colComp.SetSize(size).SetContent(sbox).GetContent()
		if stat.DeviceRole == usage.EdgeRole {
			colsEdge += cnt
		} else if stat.DeviceRole == usage.SensorRole {
			colsSensor += cnt
		}
	}

	rowEdge := components.Row().SetContent(colsEdge).GetContent()
	rowSensor := components.Row().SetContent(colsSensor).GetContent()

	boxEdge := components.Box().WithHeadBorder(true).SetHeader("Edge Devices").SetBody(rowEdge).GetContent()
	boxSensor := components.Box().WithHeadBorder(true).SetHeader("Sensor Devices").SetBody(rowSensor).GetContent()

	totalCol := colComp.SetSize(map[string]string{"md": "12"}).SetContent(boxEdge + boxSensor).GetContent()
	return components.Row().SetContent(totalCol).GetContent()
}
