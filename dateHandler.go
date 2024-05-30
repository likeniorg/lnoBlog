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
			fmt.Println(article)
			result := db.Delete(&article)
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

	mirrors := router.Group("/mirrors", isUser())
	{
		mirrors.POST("/upload_file", func(ctx *gin.Context) {
			// 处理文件上传和表单数据
			filename := ctx.PostForm("filename")
			fileComment := ctx.PostForm("fileComment")

			form, err := ctx.MultipartForm()
			if err != nil {
				ctx.String(501, "获取表单失败")
				return
			}

			files := form.File["file"]
			for _, file := range files {
				// 保存上传的文件到指定路径
				err := ctx.SaveUploadedFile(file, fmt.Sprintf("./files/%s", file.Filename))
				if err != nil {
					ctx.String(501, "文件上传失败")
					return
				}
			}

			// 可以在这里处理 filename 和 fileComment
			fmt.Printf("文件名: %s\n", filename)
			fmt.Printf("文件描述: %s\n", fileComment)

			ctx.String(200, "文件上传成功")
		})
	}
}
