package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendXML 返回xml
func SendXML(ctx *gin.Context, res string) {
	ctx.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(res))
}
