package main

import "github.com/gin-gonic/gin"

// 加载静态文件
func loadStatic(r *gin.Engine) {
	r.Static("/static", "./static")
	r.LoadHTMLGlob("web/**/*")
}
