package main

import (
	_ "k8-web-terminal/routers"
	"github.com/astaxie/beego"
	_ "k8-web-terminal/controllers"

)

func main() {
	beego.SetStaticPath("/assets", "./static/assets")
	beego.SetStaticPath("/public", "./static")
	beego.Run()
}

