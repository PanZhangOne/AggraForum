package main

import (
	"forum/bootstrap"
	"forum/conf"
	"forum/web/middleware/identity"
	"forum/web/route"
)

func main() {
	app := bootstrap.New(conf.SystemConfig.AppName, conf.SystemConfig.APPOwner)
	app.Bootstrap()
	app.Configure(identity.Configure, route.Configure)
	app.Listen(":8088")
}
