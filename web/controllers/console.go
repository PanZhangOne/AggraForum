package controllers

import (
	"forum/pkg/users"
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
	user := users.GetCurrentUser(c.Sessions)

	if user.ID == 0 {
		c.Ctx.Redirect("/login")
		return nil
	}
	results := make(map[string]interface{})

	results["User"] = user
	results["Title"] = "数据概览"
	return mvc.View{
		Layout: "shared/layout_console.html",
		Name:   "console/home/home.html",
		Data:   result.Map(results),
	}
}

// Users START
func (c *ConsoleController) GetUsersList() mvc.Result {
	var (
		user    = users.GetCurrentUser(c.Sessions)
		results = make(map[string]interface{})
	)

	results["User"] = user
	return mvc.View{
		Layout: "shared/layout_console.html",
		Name:   "console/users/list.html",
		Data:   result.Map(results),
	}
}

func (c *ConsoleController) GetUsersAdmin() mvc.Result {
	var (
		user    = users.GetCurrentUser(c.Sessions)
		results = make(map[string]interface{})
	)

	results["User"] = user
	return mvc.View{}
}

// Users END
