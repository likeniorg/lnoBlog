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

		// 删除用户
		admin.POST("/del_user", func(ctx *gin.Context) {
			user := User{}
			err := ctx.ShouldBind(&user)
			if err != nil {
				ctx.String(501, "解析参数失败")
				return
			}
			result := db.Delete(&user)
			if result.Error != nil {
				ctx.String(501, "删除用户失败")
				return
			}
			ctx.String(200, "删除用户成功")
		})

		// 新建用户前端页面
		admin.GET("/add_user.html", func(ctx *gin.Context) {

			ctx.HTML(200, "admin/add_user.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})
	}

}
