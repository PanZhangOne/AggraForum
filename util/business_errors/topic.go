package business_errors

import "github.com/kataras/iris/core/errors"

var (
	TopicTitleNotBeEmpty = errors.New("帖子标题不能为空")
)
