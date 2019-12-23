package main

import (
	"doText/router"
	"doText/system/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("hello")
	startServer()
}

func startServer(){// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// swagger 接口访问
	//if !config.IsReleaseMode() {
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
	router.UseRouters(r)

	port := config.HostPort()
	host := fmt.Sprintf(":%v", port)
	if config.IsReleaseMode() {
		//log.Infof("Listening and serviHTTP on %s\n", host)
	}
	err := r.Run(host)
	if err != nil {
		panic(errors.Wrap(err, "start server error"))
	}
}
