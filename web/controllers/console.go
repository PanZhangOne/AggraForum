package controllers

import (
	"forum/services"
	"forum/util/result"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type ConsoleController struct {
	Ctx iris.Context

	UsersService services.UsersService
	Sessions     *sessions.Session
}

func (c *ConsoleController) GetHome() mvc.Result {
	results := make(map[string]interface{})
	results["Title"] = "后台管理首页"
	return mvc.View{
		Layout: "shared/layout_console.html",
		Name:   "console/home/home.html",
		Data:   result.Map(results),
	}
}
