package route

import (
	"ZiranServer/statics"
	"ZiranServer/utils/httputils"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	gin.SetMode("release")
	r := gin.New()
	r.Use(gin.Recovery(), LoggerMiddleware())

	r.GET("/", func(context *gin.Context) {
		// context.HTML(http.StatusOK, "index.html", nil)
		httputils.SendHtmlFromEmbed(context, statics.EmbedStaticDir, "index.html")
	})

	r.GET("/index", func(context *gin.Context) {
		// context.HTML(http.StatusOK, "index.html", nil)
		httputils.SendHtmlFromEmbed(context, statics.EmbedStaticDir, "index.html")
	})

	r.GET("/index.html", func(context *gin.Context) {
		// context.HTML(http.StatusOK, "index.html", nil)
		httputils.SendHtmlFromEmbed(context, statics.EmbedStaticDir, "index.html")
	})

	//# begin-region other assets

	r.GET("/logo192.png", func(context *gin.Context) {
		httputils.SendPngFromEmbed(context, statics.EmbedStaticDir, "logo192.png")
	})

	r.GET("/logo512.png", func(context *gin.Context) {
		httputils.SendPngFromEmbed(context, statics.EmbedStaticDir, "logo512.png")
	})

	r.GET("/favicon.ico", func(context *gin.Context) {
		httputils.SendIcoFromEmbed(context, statics.EmbedStaticDir, "favicon.ico")
	})

	subStaticJS, err := fs.Sub(statics.EmbedStaticDir, "static/js")
	if err != nil {
		slog.Error(err.Error())
	}
	r.StaticFS("/static/js", http.FS(subStaticJS))

	subStaticCSS, err := fs.Sub(statics.EmbedStaticDir, "static/css")
	if err != nil {
		slog.Error(err.Error())
	}
	r.StaticFS("/static/css", http.FS(subStaticCSS))

	subStaticMedia, err := fs.Sub(statics.EmbedStaticDir, "static/media")
	if err != nil {
		slog.Error(err.Error())
	}
	r.StaticFS("/static/media", http.FS(subStaticMedia))

	r.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/")
	})

	return r
}
