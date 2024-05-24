package main

import (
	"github.com/gin-gonic/gin"
)

func adminPage(r *gin.Engine) {
	admin := r.Group("/admin", isAdmin())
	{
		// 创建博客文章版块
		admin.GET("/create_category.html", func(ctx *gin.Context) {
			ctx.HTML(200, "admin/create_category.html", gin.H{"Category": getCategory(), "authInfo": getAuthInfo(ctx)})
		})
		// 删除版块前端页面
		admin.GET("/del_category.html", func(ctx *gin.Context) {
			ctx.HTML(200, "admin/del_category.html", gin.H{"authInfo": getAuthInfo(ctx), "category": getCategory()})
		})
	}
}
