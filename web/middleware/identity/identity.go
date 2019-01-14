package identity

import (
	"forum/bootstrap"
	"github.com/kataras/iris"
	"time"
)

func New(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Sionce", time.Since(b.AppSpawnDate).String())

		ctx.Header("Server", "Golang")
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

func Configure(b *bootstrap.Bootstrapper) {
	h := New(b)
	b.UseGlobal(h)
}
