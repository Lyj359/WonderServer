package httputils

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendHtml(ctx *gin.Context, data string) {
	ctx.Data(http.StatusOK, ContentTypeHtmlUtf8, []byte(data))
}

func SendHtmlFromEmbed(ctx *gin.Context, root embed.FS, filePath string) {
	sendStaticFromEmbed(ctx, root, filePath, ContentTypeHtmlUtf8)
}

// SendJSFromEmbed 发送JavaScript数据
func SendJSFromEmbed(ctx *gin.Context, root embed.FS, filePath string) {
	sendStaticFromEmbed(ctx, root, filePath, JavascriptContentType)
}

func SendPngFromEmbed(ctx *gin.Context, root embed.FS, filePath string) {
	sendStaticFromEmbed(ctx, root, filePath, PngContentType)
}

func SendIcoFromEmbed(ctx *gin.Context, root embed.FS, filePath string) {
	sendStaticFromEmbed(ctx, root, filePath, IcoContentType)
}

func sendStaticFromEmbed(ctx *gin.Context, root embed.FS, filePath string, contentType string) {
	data, err := root.ReadFile(filePath)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, contentType, []byte(err.Error()))
	}
	ctx.Data(http.StatusOK, contentType, data)
}
