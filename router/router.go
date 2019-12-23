package router

import (
	"doText/api/test"
	"doText/system"
	"doText/system/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UseRouters 定义路由
func UseRouters(eng *gin.Engine) {
	// 首页
	eng.GET("/", index)

	hostPath := config.HostPath()
	rg := eng.Group(hostPath)

	useSystemRouters(eng, rg)
	useBizRouters(rg)
}

// 系统级通用路由
func useSystemRouters(eng *gin.Engine, rg *gin.RouterGroup) {
	rg.GET("/", system.ApiList(eng))
	rg.GET("/host", system.Host)
	rg.GET("/healthz", system.Healthz)
}

// 定义业务路由
func useBizRouters(rg *gin.RouterGroup) {
	teRg:=rg.Group("test")
	{
		teRg.POST("doTest", test.DoTest)
	}
}



// Index returns a handler
func index(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.String(
		http.StatusOK,
		"<h1>Rasse Server</h1><p>Wel come to the api server.</p><p>%v</p>",nil)
}
