package system

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Healthz 健康状况显示
func Healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}


// ApiList 显示所有路由列表
func ApiList(eng *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		routes := eng.Routes()
		total := len(routes)
		urls := map[string]string{}
		for i := 0; i < total; i++ {
			r := routes[i]
			v, exists := urls[r.Path]
			if exists {
				v = v + ", " + r.Method
			} else {
				v = r.Method
			}
			urls[r.Path] = v
		}
		ctx.JSON(http.StatusOK, urls)
	}
}

// Host 获取当前主机名
func Host(ctx *gin.Context) {
	host, _ := os.Hostname()
	ctx.String(http.StatusOK, host)
}