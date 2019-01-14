package main

import (
	"forum/bootstrap"
	"forum/conf"
	"forum/web/middleware/identity"
	"forum/web/route"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New(conf.SystemConfig.AppName, conf.SystemConfig.APPOwner)
	app.Bootstrap()
	app.Configure(identity.Configure, route.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8088")
}
