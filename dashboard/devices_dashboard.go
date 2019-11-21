package bashboard

import (
	cc "we/device_admin/dashboard/components"
	"we/device_admin/models"

	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetDashBoardContent() (types.Panel, error) {
	usageChart := cc.NewUsageChart(models.GetClusterUsage())
	deviceUnits := cc.NewDeviceUnits(models.GetClusterUsage())

	body := usageChart.GetContent() + deviceUnits.GetContent()

	return types.Panel{
		Content:     body,
		Title:       "Edge Cluster Usage",
		Description: "Edge Cluster Usage",
	}, nil
}
