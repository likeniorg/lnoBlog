package main

import (
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 公开web页面
func publicPage(router *gin.Engine) {
	// 主页
	router.GET("/", func(ctx *gin.Context) {
		// 检索文章信息并返回
		article := []Article{}
		result := db.Preload("ArticleCategory").Order("id desc").Find(&article)

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

		ctx.HTML(200, "root/index.html", gin.H{"authInfo": getAuthInfo(ctx), "articles": article, "category": getCategory()})

	})
	// 根据文章ID访问具体文章
	router.GET("article/a/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		// 文章id是数字，如果非数字停止查询并返回错误
		if _, err := strconv.Atoi(id); err != nil {
			ctx.HTML(200, "root/article.html", gin.H{"article": Article{Title: "非法参数，不正确的文章id"}, "authInfo": getAuthInfo(ctx)})
			return
		}

		// 获取文章
		article := Article{}
		result := db.Preload("ArticleCategory").Where("id = ?", id).First(&article)
		if result.Error != nil || article.Title == "" {
			ctx.HTML(200, "root/article.html", gin.H{"article": Article{Title: "该文章不存在或已被删除"}, "authInfo": getAuthInfo(ctx)})
			return
		}

		// 文章阅读数增加
		result.Update("read_count", article.ReadCount+1)

		// 标签分割
		tags := strings.Split(article.Tags, " ")

		// 获取评论
		comment := []Comment{}
		db.Preload("User").Where("article_id = ?", id).Find(&comment)

		ctx.HTML(200, "root/article.html", gin.H{"article": article, "html": template.HTML(article.ContentHTML), "authInfo": getAuthInfo(ctx), "Tags": tags, "Category": getCategory(), "Comment": comment})
	})

	// 客户端ajax获取主页版块
	r.GET("/getCategory", func(ctx *gin.Context) {
		articleCategory := []ArticleCategory{}
		result := db.Find(&articleCategory)
		if result.Error != nil {
			ctx.String(501, "服务器读取版块数据失败", nil)
			return
		}
		var htmls string
		for _, v := range articleCategory {
			html := `<h4><a href="/article/category?search=` + strconv.Itoa(int(v.ID)) + `"><span
            class="badge badge-primary">` + v.Category + `</span></a></h4>`
			htmls += html
		}
		ctx.String(200, htmls)
	})

	// 客户端ajax根据文章ID获取版块和标签
	r.GET("/getArticleCategory", func(ctx *gin.Context) {
		id := ctx.Query("id")
		// 查找版块和标签
		var article Article
		result := db.Preload("ArticleCategory").Where("id = ?", id).Select("category_id,tags").First(&article)
		if result.Error != nil {
			ctx.String(401, "不存在的文章ID")
			return
		}
		// 标签前端代码
		tagsHTML := `<hr>
		<div class="container">
			<h2>标签</h2>`
		tags := strings.Split(article.Tags, " ")
		for _, v := range tags {
			tag := `
				<h5><span class="badge badge-pill badge-primary">
						` + v + `</span></h5>`
			tagsHTML += tag
		}
		tagsHTML += "</div>"
		html := `<h4><a href="/article/category?search=` + strconv.Itoa(int(article.ArticleCategory.ID)) + `"><span
		class="badge badge-primary">` + article.ArticleCategory.Category + `</span></a></h4>`

		ctx.String(200, html+tagsHTML)
	})

	// 按照文章版块显示
	router.GET("/article/category", func(ctx *gin.Context) {
		// 文章关键字
		search := ctx.Query("search")
		// 查找指定版块的文章
		articles := []Article{}
		result := db.Preload("ArticleCategory").Where("category_id = ?", search).Find(&articles)
		_, err := strconv.Atoi(search)
		if result.Error != nil || err != nil || len(articles) == 0 {
			articles = append(articles, Article{Title: "没有找到这个版块的文章"})
		}

		ctx.HTML(200, "root/index.html", gin.H{"authInfo": getAuthInfo(ctx), "articles": articles})
	})

	// 搜索文章
	router.GET("/article/search", func(ctx *gin.Context) {
		// 文章关键字
		search := ctx.Query("search")

		// 模糊搜索文章关键字
		articles := []Article{}
		db.Find(&articles, "content_md LIKE ?", "%"+search+"%")
		if len(articles) == 0 {
			articles = append(articles, Article{Title: "没有找到这个文章"})
		}

		ctx.HTML(200, "root/index.html", gin.H{"authInfo": getAuthInfo(ctx), "articles": articles})
	})

	// 评论处理
	// 评论
	router.POST("/article/comment", func(ctx *gin.Context) {
		comment := Comment{}
		err := ctx.ShouldBind(&comment)
		if err != nil {
			ctx.String(200, err.Error())
			return
		}
		authInfo := getAuthInfo(ctx)
		if authInfo.Uid == "" {
			ctx.String(200, "请登录后评论")
		}
		uid, _ := strconv.Atoi(authInfo.Uid)
		comment.UserID = uid
		result := db.Create(&comment)
		if result.Error != nil {
			ctx.String(501, "评论创建失败 ")
			return
		} else {
			ctx.String(200, "已记录您的高见~")
		}

	})
	// 关于页面
	r.GET("/about.html", func(ctx *gin.Context) {
		ctx.HTML(200, "about/index.html", gin.H{"authInfo": getAuthInfo(ctx)})
	})

	// 下载站
	r.GET("/mirrors.html", func(ctx *gin.Context) {
		ctx.HTML(200, "mirrors/index.html", gin.H{"authInfo": getAuthInfo(ctx)})

	})

	// 管理员
	r.GET("/admin.html", func(ctx *gin.Context) {
		ctx.HTML(200, "admin/index.html", gin.H{"authInfo": getAuthInfo(ctx)})

	})
}
