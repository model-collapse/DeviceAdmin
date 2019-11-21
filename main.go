package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"

	template2 "html/template"
	"net/http"

	d "we/device_admin/dashboard"
	"we/device_admin/models"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

const modelConfigPath = "model.cfg.json"

func main() {
	cfgData, err := ioutil.ReadFile(modelConfigPath)
	if err != nil {
		log.Fatalf("Configure file does now exist...")
	}

	var mcfg models.ModelConfig
	if err := json.Unmarshal(cfgData, &mcfg); err != nil {
		log.Fatalf("Configure file has error format json...")
	}

	models.Initialize(mcfg)

	r := gin.Default()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	//template.AddLoginComp(login.GetLoginComponent())
	template.AddComp(chartjs.NewChart())

	//rootPath := "/data/www/go-admin"
	rootPath := "."

	cfg := config.ReadFromJson(rootPath + "/admin.cfg.json")
	cfg.CustomFootHtml = template2.HTML(`<div style="display:none;">
    <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
</div>`)
	cfg.CustomHeadHtml = template2.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="https://www.we.co/favicon.ico">
        <link rel="icon" type="image/png" sizes="96x96" href="https://www.we.co/favicon.ico">
        <link rel="icon" type="image/png" sizes="16x16" href="https://www.we.co/favicon.ico">`)

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return d.GetDashBoardContent()
		})
	})

	r.GET("/dashboard", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return d.GetDashBoardContent()
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
	})

	_ = r.Run(":9035")
}
