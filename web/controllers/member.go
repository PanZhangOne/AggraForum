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
	TopicService services.TopicsService
	ReplyService services.RepliesService
	Sessions     *sessions.Session
}

func (c *MemberController) GetBy(userName string) mvc.Result {
	var (
		userID        = users.GetCurrentUserID(c.Sessions)
		user          = users.GetCurrentUser(c.Sessions)
		userInfo, _   = c.UsersService.FindByUsername(userName)
		isCurrentUser = userID == userInfo.ID
		results       = make(map[string]interface{})
	)

	topics, _ := c.TopicService.FindAllNewTopicByUserID(userInfo.ID, 12, 1)
	replies := c.ReplyService.FindAllRepliesByUserID(userInfo.ID, 12, 1)

	results["User"] = user
	results["IsCurrentUser"] = isCurrentUser
	results["UserInfo"] = userInfo
	results["Topics"] = topics
	results["Replies"] = replies
	results["Title"] = userInfo.Username + " - 会员中心"

	return mvc.View{
		Name:   "member/member.html",
		Layout: "shared/layout_member.html",
		Data:   result.Map(results),
	}
}
