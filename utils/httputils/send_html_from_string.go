package httputils

import (
	"bytes"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func SendHtmlFromString(ctx *gin.Context, data string, options interface{}) {
	t, err := template.New("").Parse(data)
	if err != nil {
		slog.Error(err.Error())
		ctx.Data(http.StatusNotFound, ContentTypeHtmlUtf8, []byte(""))
		return
	}
	var out bytes.Buffer
	err = t.Execute(&out, options)
	if err != nil {
		slog.Error(err.Error())
		ctx.Data(http.StatusNotFound, ContentTypeHtmlUtf8, []byte(""))
		return
	}
	ctx.Data(http.StatusOK, ContentTypeHtmlUtf8, out.Bytes())
}
