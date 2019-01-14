package bootstrap

import (
	"forum/conf"
	"forum/util/strtime"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

func (b *Bootstrapper) SetupViews(viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true)
	// 给模板内置方法
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int64) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})

	htmlEngine.AddFunc("FromTimeString", func(t time.Time) string {
		dt := time.Unix(t.Unix(), int64(0))
		return dt.Format(conf.SysTimeform)
	})

	htmlEngine.AddFunc("FromUnixtime", func(t int64) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})

	htmlEngine.AddFunc("FromStrTime", func(t time.Time) string {
		return strtime.StrTime(t.Unix())
	})
	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:  "ssid",
		Expires: expires,
	})
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		errorMsg := ctx.Values().GetString("message")
		errorCode := ctx.GetStatusCode()
		if len(errorMsg) == 0 {
			if errorCode == 404 {
				errorMsg = "你已闯入非法领域!"
			}
		}
		err := iris.Map{
			"app":     b.AppName,
			"status":  errorCode,
			"message": errorMsg,
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.ViewLayout("shared/layout_error.html")
		_ = ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./web/public/"
	//Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./web/views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb("/public", StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	_ = b.Run(iris.Addr(addr), cfgs...)
}
