package main

import (
	"html/template"

    tmpl "github.com/atEaE/aws-file-trigger-sysexam/web/templates"
    router "github.com/atEaE/aws-file-trigger-sysexam/web/config/routers"
)

func init() {
	tmplConf := tmpl.TmplConfig{
		Root:         "views",
		Extension:    ".html",
		Master:       "shared/_layout",
		Partials:     []string{},
		Functions:    make(template.FuncMap),
		DisableCache: false,
		Delmis:       tmpl.Delims{Left: "{{", Right: "}}"},
    }
    
    router.Echo.Renderer = tmpl.New(tmplConf)
    router.Echo.Static("public", "public")
}

func main() {
    router.Echo.Start(":8080")
}
