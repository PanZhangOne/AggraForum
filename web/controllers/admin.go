package controllers

import (
	"forum/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type AdminController struct {
	Ctx iris.Context

	UsersService services.UsersService
	Sessions     *sessions.Session
}


