package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DoTest(ctx *gin.Context){
	ctx.JSON(http.StatusOK,0)
}