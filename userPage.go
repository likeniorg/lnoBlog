package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 用户操作页面
func userPage(router *gin.Engine) {
	// 文章前端页面，包括了我的文章、创建文章、修改文章、删除文章
	article := router.Group("/article", isUser())
	{
		//我的文章
		article.GET("/my_article.html", func(ctx *gin.Context) {
			// 获取客户端身份
			authInfo := getAuthInfo(ctx)

			// 检索文章信息并返回
			article := []Article{}
			result := db.Where("author_id = ?", authInfo.Uid).Find(&article)

			// 如果数据库读取出错或者文章数量为0
			if result.Error != nil || len(article) == 0 {
				var errMessage string
				if result.Error != nil {
					errMessage = "从数据库中读取文章失败"
				} else {
					errMessage = "欢迎访客物联网知识在线问答社区"
				}
				article = append(article, Article{Title: errMessage, Model: Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}})
			}

			ctx.HTML(200, "root/my_article.html", gin.H{"authInfo": getAuthInfo(ctx), "articles": article})
		})

		// 创建文章
		article.GET("/create_article.html", func(ctx *gin.Context) {
			ctx.HTML(200, "root/create_article.html", gin.H{"Category": getCategory(), "authInfo": getAuthInfo(ctx)})
		})

		// 修改文章
		article.GET("/update_article.html", func(ctx *gin.Context) {
			article := []Article{}
			result := db.Select("id,title").Find(&article)
			if result.Error != nil {
				article = append(article, Article{Title: "查找文章错误"})
			}
			ctx.HTML(200, "root/update_article.html", gin.H{"Category": getCategory(), "authInfo": getAuthInfo(ctx), "articles": article})
		})

		// 删除文章
		article.GET("/del_article.html", func(ctx *gin.Context) {
			// 检索文章信息并返回
			article := []Article{}
			result := db.Select("id,title").Order("id desc").Find(&article)

			// 如果数据库读取出错或者文章数量为0
			if result.Error != nil || len(article) == 0 {
				var errMessage string
				if result.Error != nil {
					looger.Error("从数据库中读取文章失败")
					errMessage = "从数据库中读取文章失败"
				} else {
					errMessage = "还未创建文章~"
				}
				article = append(article, Article{Title: errMessage, Model: Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}})
			}

			ctx.HTML(200, "root/del_article.html", gin.H{"authInfo": getAuthInfo(ctx), "article": article})
		})

		// 我的评论
		article.GET("/my_comment.html", func(ctx *gin.Context) {
			authInfo := getAuthInfo(ctx)
			// 获取评论
			comment := []Comment{}

			result := db.Preload("User").Preload("Article").Where("user_id = ?", authInfo.Uid).Find(&comment)
			if result.Error != nil || result.RowsAffected == 0 {
				comment = append(comment, Comment{Model: Model{ID: 0}, CommentText: "你还没有发表高见~"})
			}
			ctx.HTML(200, "root/my_comment.html", gin.H{"authInfo": authInfo, "Category": getCategory(), "Comment": comment})
		})

		// 归属于修改文章前端页面，通过ajax返回指定的文章数据等待修改
		article.GET("/getArticle", func(ctx *gin.Context) {
			id := ctx.Query("id")
			article := Article{}
			db.Preload("ArticleCategory").Where("id = ?", id).Find(&article)
			ctx.JSON(200, gin.H{"article": article})
		})
	}

	// 在线问答前端页面，包含了我的提问、发起提问、删除提问
	problem := router.Group("/problem", isUser())
	{
		//我的提问
		problem.GET("/my_problem.html", func(ctx *gin.Context) {
			// 获取客户端身份
			authInfo := getAuthInfo(ctx)
			// 检索文章信息并返回
			article := []Article{}
			result := db.Where("author_id = ? AND category_id = ?", authInfo.Uid, 2).Find(&article)
			// 如果数据库读取出错或者文章数量为0
			if result.Error != nil || len(article) == 0 {
				var errMessage string
				if result.Error != nil {
					errMessage = "从数据库中读取问题失败"
				} else {
					errMessage = "欢迎访客物联网知识在线问答社区"
				}
				article = append(article, Article{Title: errMessage, Model: Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}})
			}
			ctx.HTML(200, "root/my_problem.html", gin.H{"authInfo": getAuthInfo(ctx), "articles": article})
		})

		// 发起问题前端页面，后台处理用的是/article/create_article路径，他们的属于同一个数据表，区别在于版块不同
		problem.GET("/create_problem.html", func(ctx *gin.Context) {
			ctx.HTML(200, "root/create_problem.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})

		// 删除提问前端页面
		problem.GET("/del_problem.html", func(ctx *gin.Context) {
			// 检索文章信息并返回
			article := []Article{}
			result := db.Select("id,title,category_id").Where("category_id = ?", 2).Order("id desc").Find(&article)

			// 如果数据库读取出错或者文章数量为0
			if result.Error != nil || len(article) == 0 {
				var errMessage string
				if result.Error != nil {
					errMessage = "从数据库中读取文章失败"
				} else {
					errMessage = "还未创建文章~"
				}
				article = append(article, Article{Title: errMessage, Model: Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}})
			}
			ctx.HTML(200, "root/del_problem.html", gin.H{"authInfo": getAuthInfo(ctx), "article": article})
		})
	}

	// 修改用户名、密码、邮箱前端页面
	user := router.Group("/user", isUser())
	{
		// 修改用户名
		user.GET("/update_user.html", func(ctx *gin.Context) {
			ctx.HTML(200, "root/update_user.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})

		// 修改密码
		user.GET("/update_pass.html", func(ctx *gin.Context) {
			ctx.HTML(200, "root/update_pass.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})

		//修改邮箱
		user.GET("/update_email.html", func(ctx *gin.Context) {
			ctx.HTML(200, "root/update_email.html", gin.H{"authInfo": getAuthInfo(ctx)})
		})
	}

}
