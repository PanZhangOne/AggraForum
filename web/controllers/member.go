package controllers

import (
	"forum/services"
	"forum/util/result"
	"forum/util/users"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type MemberController struct {
	Ctx iris.Context

	UsersService services.UsersService
	Sessions     *sessions.Session
}

func (c *MemberController) GetBy(userName string) mvc.Result {
	var (
		userID        = users.GetCurrentUserID(c.Sessions)
		user          = users.GetCurrentUser(c.Sessions)
		userInfo, _   = c.UsersService.FindByUsername(userName)
		results       = make(map[string]interface{})
		isCurrentUser = userID == userInfo.ID
	)

	results["User"] = user
	results["IsCurrentUser"] = isCurrentUser
	results["UserInfo"] = userInfo
	results["Title"] = "会员中心"

	return mvc.View{
		Name:   "member/member.html",
		Layout: "shared/layout_member.html",
		Data:   result.Map(results),
	}
}
