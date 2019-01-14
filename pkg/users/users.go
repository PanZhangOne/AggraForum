package users

import (
	"forum/conf"
	"forum/entitys"
	"forum/services"
	"github.com/kataras/iris/sessions"
)

var userService = services.NewUserService()

// 获取当前登录的用户ID
func GetCurrentUserID(sess *sessions.Session) uint {
	userID := sess.GetIntDefault(conf.SystemConfig.UserIDKey, 0)
	return uint(userID)
}

// 获取当前登录的用户
func GetCurrentUser(sess *sessions.Session) *entitys.User {
	user, _ := userService.FindByID(GetCurrentUserID(sess))
	return user
}
