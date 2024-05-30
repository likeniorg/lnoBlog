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

		admin.GET("/all_user.html", func(ctx *gin.Context) {
			users := []User{}
			result := db.Select("id,username,email").Find(&users)
			if result.Error != nil {
				ctx.String(501, "查找用户数据失败")
				return
			}
			ctx.HTML(200, "admin/all_user.html", gin.H{"authInfo": getAuthInfo(ctx), "users": users})
		})

		// 删除用户前端页面
		admin.GET("/del_user.html", func(ctx *gin.Context) {
			users := []User{}
			ctx.ShouldBind(&users)
			result := db.Find(&users)
			if result.Error != nil {
				users = append(users, User{Username: "查找用户数据失败"})
			}
			ctx.HTML(200, "admin/del_user.html", gin.H{"authInfo": getAuthInfo(ctx), "users": users})
		})

		// 新建用户前端页面
		admin.GET("/add_user.html", func(ctx *gin.Context) {

			ctx.HTML(200, "admin/add_user.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})

		// 可以发送给任意邮箱
		admin.GET("/send_email.html", func(ctx *gin.Context) {

			ctx.HTML(200, "admin/send_email.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})

		// 指定用户邮箱号发送邮件
		admin.GET("/send_smail.html", func(ctx *gin.Context) {

			ctx.HTML(200, "admin/send_smail.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})
	}

}
