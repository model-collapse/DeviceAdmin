package devicebox

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/themes/adminlte/components"
)

type DeviceBox struct {
	components.Base
	Title        template.HTML
	Value        template.HTML
	Url          string
	Color        template.HTML
	IsAwsomeIcon bool
	IsHexColor   bool
	Icon         template.HTML
}

func New() DeviceBox {
	return DeviceBox{}
}

func (s DeviceBox) SetTitle(value template.HTML) DeviceBox {
	s.Title = value
	return s
}

func (s DeviceBox) SetValue(value template.HTML) DeviceBox {
	s.Value = value
	return s
}

func (s DeviceBox) SetColor(value template.HTML) DeviceBox {
	s.Color = value
	if strings.Contains(string(value), "#") {
		s.IsHexColor = true
	}
	return s
}

func (s DeviceBox) SetIcon(value template.HTML) DeviceBox {
	s.Icon = value
	if !strings.Contains(string(value), "<") {
		s.IsAwsomeIcon = true
	}
	return s
}

func (s DeviceBox) SetUrl(value string) DeviceBox {
	s.Url = value
	return s
}

func (s DeviceBox) GetTemplate() (*template.Template, string) {
	ctype := "custom_icon"
	if s.IsAwsomeIcon {
		ctype = "awsome_icon"
	}

	tmpl, err := template.New("devicebox").
		Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
			"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
				if cdnUrl == "" {
					return prefixUrl + assetsUrl
				}
				return cdnUrl + assetsUrl
			},
			"isLinkUrl": func(s string) bool {
				return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
			},
		}).
		Parse(List[ctype])

	if err != nil {
		logger.Error("DeviceBox GetTemplate Error: ", err)
	}

	return tmpl, "devicebox"
}

func (s DeviceBox) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := s.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, s)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
