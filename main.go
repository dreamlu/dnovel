package main

import (
	"dnovel/routers"
	"github.com/dreamlu/gt/tool/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	//log.Println(gt.Version)
	gin.SetMode(gin.DebugMode)
	//r := routers.SetRouter()
	// 性能调试
	//pprof.Register(routers.Router)
	// Listen and Server in 0.0.0.0:8080
	_ = routers.Router.Run(":" + conf.GetString("app.port"))
}
