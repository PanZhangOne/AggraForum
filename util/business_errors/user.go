package business_errors

import "github.com/kataras/iris/core/errors"

var (
	// Username
	UsernameAlreadyExists = errors.New("用户名已存在")
	UsernameNotExist      = errors.New("用户名不存在")
	UsernameNotBeEmpty    = errors.New("用户名不能为空")

	// Email
	EmailAlreadyExists = errors.New("该邮箱已存在")
	EmailNotExist      = errors.New("该邮箱不存在")
	EmailNotBeEmpty    = errors.New("邮箱不能为空")

	// Password Password must not be less than eight characters
	PasswordNotBeEmpty              = errors.New("密码不能为空")
	PasswordError                   = errors.New("密码错误")
	PasswordLessThanEightCharacters = errors.New("密码不能少于八个字符")

	// Accounts
	AccountFrozen = errors.New("账户被冻结")
)
