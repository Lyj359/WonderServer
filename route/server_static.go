package route

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerStatic(prefix string, embedFs embed.FS) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 去掉前缀
		fsys, err := fs.Sub(embedFs, prefix)
		if err != nil {
			panic(err)
		}
		fs2 := http.FS(fsys)
		f, err := fs2.Open(ctx.Request.URL.Path)
		if err != nil {
			// 判断文件不存在，退出交给其他路由函数
			ctx.Next()
			return
		}
		defer f.Close()
		http.FileServer(fs2).ServeHTTP(ctx.Writer, ctx.Request)
		ctx.Abort()
	}
}
