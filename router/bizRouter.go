package router

import (
	"doText/api/test"
	"github.com/gin-gonic/gin"
)

// 定义业务路由
func useBizRouters(rg *gin.RouterGroup) {
	teRg:=rg.Group("test")
	{
		teRg.POST("doTest", test.DoTest)
	}
}
