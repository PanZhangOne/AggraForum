package controllers

import (
	"forum/pkg/users"
	"forum/services"
	"forum/util/result"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type MemberController struct {
	Ctx iris.Context

	UsersService services.UsersService
	Sessions     *sessions.Session
}

func (c *MemberController) Get() mvc.Result {
	var (
		user    = users.GetCurrentUser(c.Sessions)
		results = make(map[string]interface{})
	)

	results["User"] = user

	return mvc.View{
		Name:   "member/member.html",
		Layout: "shared/layout_member.html",
		Data:   result.Map(results),
	}
}

func (c *MemberController) GetBy(userName string) mvc.Result {
	var (
		userID        = users.GetCurrentUserID(c.Sessions)
		user          = users.GetCurrentUser(c.Sessions)
		userInfo, _   = c.UsersService.FindByUsername(userName)
		results       = make(map[string]interface{})
		isCurrentUser = userID == userInfo.ID
	)

	if isCurrentUser {
		c.Ctx.Redirect("/member")
		return nil
	}

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
