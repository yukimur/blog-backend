package main

import (
	_ "Blog/routers"

	beego "github.com/beego/beego/v2/server/web"
	_ "Blog/models"
)


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
