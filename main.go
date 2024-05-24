package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var r = gin.New()

// 数据库
var db *gorm.DB

func main() {
	getConf()
	db = conn()

	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// 你的自定义格式
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 		param.Path,
	// 	)
	// }))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(authorization())
	// 加载静态资源
	loadStatic(r)

	// 加载路由
	router(r)
	userHandler()

	// 如果域名等于空，处于测试环境，否则部署为生产环境
	if Domain == "" {
		r.Run(":8080")
	} else {
		go r.RunTLS(":443", "ssl/"+Domain+".crt", "ssl/"+Domain+".key")
		r.Run(":80")
	}

}
