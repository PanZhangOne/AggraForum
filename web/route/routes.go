package route

import (
	"forum/bootstrap"
	"forum/services"
	"forum/web/controllers"
	"github.com/kataras/iris/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	var (
		userService         = services.NewUserService()
		labelService        = services.NewLabelService()
		topicService        = services.NewTopicsService()
		repliesService      = services.NewRepliesService()
		likeTopicService    = services.NewLikeTopicService()
		collectTopicService = services.NewCollectTopicService()
		messageService      = services.NewMessageService()
	)

	client := mvc.New(b.Party("/"))
	client.Register(userService, labelService, topicService, repliesService,
		collectTopicService, likeTopicService, messageService,
		b.Sessions.Start)
	client.Handle(new(controllers.ClientController))

	member := mvc.New(b.Party("/member"))
	member.Register(userService, labelService, topicService, repliesService,
		collectTopicService, likeTopicService, b.Sessions.Start)
	member.Handle(new(controllers.MemberController))

	console := mvc.New(b.Party("/console"))
	console.Register(userService, labelService, topicService, repliesService,
		collectTopicService, likeTopicService, b.Sessions.Start)
	console.Handle(new(controllers.ConsoleController))
}
