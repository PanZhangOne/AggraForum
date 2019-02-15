package controllers

import (
	"fmt"
	"forum/conf"
	"forum/entitys"
	"forum/services"
	"forum/util/avatar"
	"forum/util/result"
	"forum/util/users"
	"github.com/iris-contrib/blackfriday"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
	"strconv"
)

type ClientController struct {
	Ctx iris.Context

	UsersService        services.UsersService
	LabelService        services.LabelsService
	TopicService        services.TopicsService
	RepliesService      services.RepliesService
	CollectTopicService services.CollectTopicService
	LikeTopicService    services.LikeTopicService
	MessageService      services.MessageService

	Sessions *sessions.Session
}

// loginOut sign out
func (c *ClientController) loginOut() {
	c.Sessions.Destroy()
}

func (c *ClientController) Get() mvc.Result {

	hots := c.TopicService.FindHots(10)
	hotLabels := c.LabelService.FindHotLabels()
	user := users.GetCurrentUser(c.Sessions)
	topics, _ := c.TopicService.FindAllNewsTopics()

	var results = make(map[string]interface{})
	results["Title"] = "首页"
	results["User"] = user
	results["Hots"] = hots
	results["HotLabels"] = hotLabels
	results["Topics"] = topics
	return mvc.View{
		Name: "index.html",
		Data: result.Map(results),
	}
}

func (c *ClientController) GetLogin() mvc.Result {
	c.loginOut()
	var (
		results   = make(map[string]interface{})
		hots      = c.TopicService.FindHots(10)
		hotLabels = c.LabelService.FindHotLabels()
	)

	results["Hots"] = hots
	results["HotLabels"] = hotLabels
	results["Title"] = "登录"

	return mvc.View{
		Name: "login.html",
		Data: result.Map(results),
	}
}

/**
 * @api {post} /login
 */
func (c *ClientController) PostLogin() {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
		results  = make(map[string]interface{})
	)

	user, err := c.UsersService.Login(username, password)

	if err != nil {
		results["success"] = false
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	if user.ID <= 0 {
		results["success"] = false
		results["message_status"] = "登录失败"
		_, _ = c.Ctx.JSON(results)
	}

	c.Sessions.Set(conf.SystemConfig.UserIDKey, int(user.ID))
	results["success"] = true
	results["message_status"] = "登录成功"
	results["return_url"] = "/"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) AnyLogout() {
	c.loginOut()
	c.Ctx.Redirect("/login")
}

func (c *ClientController) GetSignup() mvc.Result {
	var (
		results   = make(map[string]interface{})
		hots      = c.TopicService.FindHots(10)
		hotLabels = c.LabelService.FindHotLabels()
	)

	results["Hots"] = hots
	results["HotLabels"] = hotLabels
	results["Title"] = "注册"

	return mvc.View{
		Name: "signup.html",
		Data: result.Map(results),
	}
}

// PostSignup
func (c *ClientController) PostSignup() {
	results := make(map[string]interface{})
	user := entitys.User{}
	username := c.Ctx.FormValue("username")
	password := c.Ctx.FormValue("password")
	email := c.Ctx.FormValue("email")

	status := true
	message := ""

	if len(username) == 0 {
		status = false
		message = "用户名不能为空!"
	}
	if len(password) == 0 {
		status = false
		message = "密码不能为空!"
	}
	if len(email) == 0 {
		status = false
		message = "邮箱不能为空!"
	}

	if !status {
		results["success"] = status
		results["message_status"] = message
		_, _ = c.Ctx.JSON(results)
		return
	}

	// 生成用户头像
	avatar.GenerateAvatarFromUsername(username)
	user.Username = username
	user.Password = password
	user.Email = email
	user.Avatar = "/public/avatar/" + username + ".png"

	err := c.UsersService.Create(&user)
	if err != nil {
		status = false
		message = err.Error()
		results["success"] = status
		results["message_status"] = message
		_, _ = c.Ctx.JSON(results)
		return
	}

	results["success"] = status
	results["message_status"] = "注册成功"
	results["return_url"] = "/login"
	_, _ = c.Ctx.JSON(results)
	return
}

func (c *ClientController) GetNew() mvc.Result {
	userID := users.GetCurrentUserID(c.Sessions)
	if userID <= 0 {
		c.Ctx.Redirect("/login")
	}

	var (
		results   = make(map[string]interface{})
		user, _   = c.UsersService.FindByID(userID)
		hots      = c.TopicService.FindHots(10)
		hotLabels = c.LabelService.FindHotLabels()
		labels    = c.LabelService.FindAllLabel()
	)

	results["User"] = user
	results["Hots"] = hots
	results["HotLabels"] = hotLabels
	results["Labels"] = labels
	results["New"] = true
	results["Title"] = "发表新主题"
	results["Editor"] = true
	return mvc.View{
		Name: "topic/new.html",
		Data: result.Map(results),
	}
}

func (c *ClientController) PostNew() {
	userID := users.GetCurrentUserID(c.Sessions)

	if userID <= 0 {
		c.Ctx.Redirect("/login")
		return
	}

	var (
		title   = c.Ctx.FormValue("title")
		labelID = c.Ctx.FormValue("label_id")
		content = c.Ctx.FormValue("content")
		results = make(map[string]interface{})
	)

	if len(title) <= 0 {
		results["success"] = false
		results["message_status"] = "主题标题不能为空"
		_, _ = c.Ctx.JSON(results)
		return
	}
	_labelID, _ := strconv.Atoi(labelID)

	topic := entitys.Topic{}
	topic.Title = title
	topic.UserId = userID
	topic.LabelId = uint(_labelID)
	topic.Content = content
	err := c.TopicService.Create(&topic)

	if err != nil {
		results["success"] = false
		results["message_status"] = err.Error()
		fmt.Println(err)
		_, _ = c.Ctx.JSON(results)
		return
	}

	results["success"] = true
	results["message_status"] = "发表成功"
	results["return_url"] = "/label/" + labelID
	_, _ = c.Ctx.JSON(results)
	return
}

func (c *ClientController) GetLabelBy(id uint) mvc.Result {
	var (
		results   = make(map[string]interface{})
		hots      = c.TopicService.FindHots(10)
		hotLabels = c.LabelService.FindHotLabels()
	)

	label, _ := c.LabelService.FindByID(id)
	topics, _ := c.TopicService.FindAllByLabelID(id, 50, 0)
	user := users.GetCurrentUser(c.Sessions)

	results["User"] = user
	results["Topics"] = topics
	results["Title"] = label.LabelName
	results["Label"] = label
	results["HotLabels"] = hotLabels
	results["Hots"] = hots
	return mvc.View{
		Name: "label/label.html",
		Data: result.Map(results),
	}
}

func (c *ClientController) GetTopicBy(id uint) mvc.Result {
	var (
		user             = users.GetCurrentUser(c.Sessions)
		topic, _         = c.TopicService.FindByID(id)
		hotLabels        = c.LabelService.FindHotLabels()
		results          = make(map[string]interface{})
		replies          = c.RepliesService.FindRepliesByTopicID(id)
		hots             = c.TopicService.FindHots(10)
		isCollect        = c.CollectTopicService.CheckCollectedTopic(user.ID, topic.ID)
		likeOrDislike, _ = c.LikeTopicService.FindTopicIsLikeOrDislike(user.ID, id)
	)

	if likeOrDislike.ID >= 1 {
		topic.Like = likeOrDislike.Like
		topic.Dislike = likeOrDislike.Dislike
	}

	unsafe := blackfriday.Run([]byte(topic.Content))
	contentHtml := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	for idx, reply := range replies {
		us := blackfriday.Run([]byte(reply.Content))
		replies[idx].ContentHtml = template.HTML(bluemonday.UGCPolicy().SanitizeBytes(us))
	}

	results["User"] = user
	results["Topic"] = topic
	results["Title"] = topic.Title
	results["HotLabels"] = hotLabels
	results["Content"] = contentHtml
	results["Replies"] = replies
	results["Hots"] = hots
	results["RepliesLen"] = len(replies)
	results["Editor"] = true
	results["IsCollect"] = isCollect

	return mvc.View{
		Name: "topic/topic",
		Data: result.Map(results),
	}
}

func (c *ClientController) GetHots() mvc.Result {
	var (
		user      = users.GetCurrentUser(c.Sessions)
		topics    = c.TopicService.FindHots(50)
		hotLabels = c.LabelService.FindHotLabels()
		results   = make(map[string]interface{})
		hots      = c.TopicService.FindHots(10)
	)

	results["User"] = user
	results["HotLabels"] = hotLabels
	results["Topics"] = topics
	results["Hots"] = hots
	results["Title"] = "社区最热"
	return mvc.View{
		Name: "hots.html",
		Data: result.Map(results),
	}
}

func (c *ClientController) PostReply() {
	userID := users.GetCurrentUserID(c.Sessions)
	if userID == 0 {
		c.Ctx.Redirect("/login")
		return
	}
	var (
		topicID    = c.Ctx.FormValue("topic_id")
		content    = c.Ctx.FormValue("content")
		parentID   = c.Ctx.FormValue("parent_id")
		deviceInfo = c.Ctx.FormValue("device_info")
		results    = make(map[string]interface{})
	)
	results["success"] = false
	results["message_status"] = ""

	var reply = new(entitys.Reply)
	_topicID, err := strconv.Atoi(topicID)
	if err != nil {
		_topicID = 0
	}
	_parentID, err := strconv.Atoi(parentID)
	if err != nil {
		_parentID = 0
	}
	user, err := c.UsersService.FindByID(userID)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	reply.UserID = user.ID
	reply.TopicID = uint(_topicID)
	reply.ParentID = uint(_parentID)
	reply.Content = content
	reply.DeviceInfo = deviceInfo

	err = c.RepliesService.Reply(reply)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = true
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetNode() mvc.Result {
	var (
		user      = users.GetCurrentUser(c.Sessions)
		hotLabels = c.LabelService.FindHotLabels()
		labels    = c.LabelService.FindAllLabel()
		results   = make(map[string]interface{})
		hots      = c.TopicService.FindHots(10)
	)

	results["User"] = user
	results["Title"] = "节点"
	results["HotLabels"] = hotLabels
	results["Labels"] = labels
	results["Hots"] = hots

	return mvc.View{
		Name: "node.html",
		Data: result.Map(results),
	}
}

func (c *ClientController) GetCollectTopicBy(id uint) {
	var (
		user    = users.GetCurrentUser(c.Sessions)
		results = make(map[string]interface{})
	)
	results["success"] = false
	results["message_status"] = ""

	if user.ID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	topic, _ := c.TopicService.FindByID(id)
	if topic.ID <= 0 {
		results["message_status"] = "未找到该主题"
		_, _ = c.Ctx.JSON(results)
		return
	}

	err := c.CollectTopicService.Collect(user.ID, topic.ID, topic.LabelId)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = true
	results["message_status"] = "收藏成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetCollectTopicCancelBy(topicID uint) {
	var (
		user    = users.GetCurrentUser(c.Sessions)
		results = make(map[string]interface{})
	)
	results["success"] = false
	results["message_status"] = ""

	if user.ID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	topic, _ := c.TopicService.FindByID(topicID)
	if topic.ID <= 0 {
		results["message_status"] = "未找到该主题"
		_, _ = c.Ctx.JSON(results)
		return
	}

	err := c.CollectTopicService.UnCollect(user.ID, topicID)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = true
	results["message_status"] = "取消收藏成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetTopicLikeBy(topicID uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	results["success"] = false
	results["message_status"] = ""

	if userID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	ok, err := c.LikeTopicService.Like(userID, topicID)

	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = ok
	results["message_status"] = "点赞成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetTopicLikeCancelBy(topicID uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	results["success"] = false
	results["message_status"] = ""

	if userID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	ok, err := c.LikeTopicService.CancelLike(userID, topicID)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = ok
	results["message_status"] = "取消点赞成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetTopicDislikeBy(topicID uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	if userID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	ok, err := c.LikeTopicService.Dislike(userID, topicID)
	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = ok
	results["message_status"] = "不喜欢成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetTopicDislikeCancelBy(topicID uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	if userID <= 0 {
		results["message_status"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}
	ok, err := c.LikeTopicService.CancelDislike(userID, topicID)

	if err != nil {
		results["message_status"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = ok
	results["message_status"] = "取消不喜欢成功"
	_, _ = c.Ctx.JSON(results)
}

func (c *ClientController) GetSetting() mvc.Result {
	var (
		user       = users.GetCurrentUser(c.Sessions)
		topicCount = c.TopicService.GetTopicCount(user.ID)
		results    = make(map[string]interface{})
	)

	if user.ID <= 0 {
		c.Ctx.Redirect("/login")
		return nil
	}

	results["User"] = user
	results["Title"] = "个人设置"
	results["TopicCounts"] = topicCount

	return mvc.View{
		Name:   "setting.html",
		Layout: "shared/layout_member.html",
		Data:   result.Map(results),
	}
}

func (c *ClientController) GetMessage() mvc.Result {
	var (
		user              = users.GetCurrentUser(c.Sessions)
		topicCount        = c.TopicService.GetTopicCount(user.ID)
		results           = make(map[string]interface{})
		messages, _       = c.MessageService.GetAllMessagesByUser(user.ID, 12, 1)
		unreadMessages, _ = c.MessageService.GetAllNotReadMessages(user.ID, 12, 1)
	)

	if user.ID <= 0 {
		c.Ctx.Redirect("/login")
		return nil
	}

	results["User"] = user
	results["Title"] = "消息中心"
	results["TopicCounts"] = topicCount
	results["Messages"] = messages
	results["UnReadMessages"] = unreadMessages

	return mvc.View{
		Name:   "message.html",
		Layout: "shared/layout_member.html",
		Data:   result.Map(results),
	}
}

func (c *ClientController) GetMessageRead(id uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	results["success"] = false
	results["message"] = ""

	if userID <= 0 {
		results["message"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	err := c.MessageService.ReadMessage(userID, id)
	if err != nil {
		results["message"] = err.Error()
		_, _ = c.Ctx.JSON(results)
		return
	}
	results["success"] = true
	results["message"] = "success"
	_, _ = c.Ctx.JSON(results)
	return
}

func (c *ClientController) GetMessageDelete(id uint) {
	var (
		userID  = users.GetCurrentUserID(c.Sessions)
		results = make(map[string]interface{})
	)

	results["success"] = false
	results["message"] = ""

	if userID <= 0 {
		results["message"] = "请先登录"
		_, _ = c.Ctx.JSON(results)
		return
	}

	err := c.MessageService.DeleteMessage(userID, id)
	if err != nil {
		results["message"] = err.Error()
		return
	}
	results["success"] = true
	results["message"] = "删除成功"
	_, _ = c.Ctx.JSON(results)
	return
}

func (c *ClientController) GetAbout() mvc.Result {
	hots := c.TopicService.FindHots(10)
	hotLabels := c.LabelService.FindHotLabels()
	user := users.GetCurrentUser(c.Sessions)
	topics, _ := c.TopicService.FindAllNewsTopics()

	var results = make(map[string]interface{})

	results["Title"] = "关于本站"
	results["User"] = user
	results["Hots"] = hots
	results["HotLabels"] = hotLabels
	results["Topics"] = topics
	return mvc.View{
		Name: "about.html",
		Data: result.Map(results),
	}
}
