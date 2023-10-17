package main

import (
	_ "learn_go/routers"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/aplikasi.log","level":7,"maxlines":0,"maxsize":1000000,"daily":true,"maxdays":10,"color":true}`)

	if err != nil {
		log.Fatal(err)
	}

	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))
	beego.Run()
}
