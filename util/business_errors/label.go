package business_errors

import "github.com/kataras/iris/core/errors"

var (
	LabelNameAlreadyExists = errors.New("用户名已存在")
	LabelNameNotExist      = errors.New("用户名不存在")
	LabelNameNotBeEmpty    = errors.New("用户名不能为空")
)
