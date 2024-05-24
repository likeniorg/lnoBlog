package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func dataHandler(router *gin.Engine) {

	// 文章后台操作处理
	article := router.Group("/article", isUser())
	{
		// 创建文章后台处理
		article.POST("/create_article", func(ctx *gin.Context) {
			var article = Article{}
			err := ctx.ShouldBind(&article)
			if err != nil {
				ctx.String(501, "参数解析失败")
				return
			}
			article.ContentHTML, err = markdownToHtmlString(article.ContentMD)
			if err != nil {
				fmt.Println(err)
				ctx.String(501, "转换为markdown失败")
				return
			}

			article.CreatedAt = time.Now()

			cookie, err := ctx.Cookie("authorization")
			if err != nil {
				ctx.String(401, "客户端Cookie不正确,请尝试重新登录")
				return
			}
			claims, err := ParseJwt(cookie)
			if err != nil {
				ctx.String(401, "身份验证失败，请重新登录")
				return
			}

			id, _ := strconv.Atoi(claims.ID)
			article.AuthorID = id
			db.Create(&article)

			ctx.String(200, "创建文章成功~")
		})

		// 修改文章后台处理
		article.POST("/update_article", func(ctx *gin.Context) {
			var article = Article{}
			err := ctx.ShouldBind(&article)
			if err != nil {
				ctx.String(501, "参数解析失败")
				return
			}
			article.ContentHTML, err = markdownToHtmlString(article.ContentMD)
			if err != nil {
				fmt.Println(err)
				ctx.String(501, "转换为markdown失败")
				return
			}

			article.UpdatedAt = time.Now()

			cookie, err := ctx.Cookie("authorization")
			if err != nil {
				ctx.String(401, "客户端Cookie不正确,请尝试重新登录")
				return
			}
			claims, err := ParseJwt(cookie)
			if err != nil {
				ctx.String(401, "身份验证失败，请重新登录")
				return
			}

			id, _ := strconv.Atoi(claims.ID)
			article.AuthorID = id
			result := db.Model(&article).Updates(article)

			if result.Error != nil {
				ctx.String(501, result.Error.Error())
				return
			}
			ctx.String(200, "修改文章成功~")
		})

		// 创建版块
		article.POST("/create_category", func(ctx *gin.Context) {
			var category = ArticleCategory{}
			err := ctx.ShouldBind(&category)
			if err != nil {
				fmt.Println(err)
			}

			result := db.Create(&category)

			fmt.Println(result)
			ctx.String(200, "版块创建成功")
		})

		// 删除文章
		article.POST("/del_article", func(ctx *gin.Context) {
			article := Article{}
			ctx.ShouldBind(&article)
			result := db.Where("id = ?", article.ID).Delete(&article)
			if result.Error != nil {
				ctx.String(500, "删除文章失败")
				return
			}
			ctx.String(200, "删除文章成功")

		})

		// 删除版块
		article.POST("/del_category", func(ctx *gin.Context) {
			category := ArticleCategory{}
			ctx.ShouldBind(&category)
			result := db.Where("id = ?", category.ID).Delete(&category)
			if result.Error != nil {
				ctx.String(500, "删除版块失败，如果该版块还存在文章无法删除!")
				return
			}
			ctx.String(200, "删除版块成功")
		})

	}

}
