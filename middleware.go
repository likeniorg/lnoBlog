package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// 身份验证
func authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identify := []string{"nologin", "user", "admin"}
		cookie, err := ctx.Cookie("authorization")
		if err != nil {
			ctx.Set("AuthInfo", AuthInfo{Identify: identify[0]})
			ctx.Next()
			return
		}
		claims, err := ParseJwt(cookie)
		if err != nil {
			looger.Warn("错误的Token", "错误Token", cookie)
			ctx.Set("AuthInfo", AuthInfo{Identify: identify[0]})
			ctx.Next()
			return
		}

		if claims.Identify == identify[1] || claims.Identify == identify[2] {
			ctx.Set("AuthInfo", AuthInfo{Identify: claims.Identify, Username: claims.Username, Uid: claims.ID})
			ctx.Next()
			return
		}
		logWarn("不属于已指定身份访问")
		ctx.Next()
	}
}

// 判断用户是否登录
func isUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authInfo := getAuthInfo(ctx)
		if authInfo.Identify == "user" || authInfo.Identify == "admin" {
			ctx.Next()
			return
		}
		articles := []Article{}
		articles = append(articles, Article{Title: "没有访问权限，请登录", Model: Model{CreatedAt: time.Now()}})
		ctx.Abort()
		ctx.HTML(200, "root/403.html", gin.H{"authInfo": authInfo, "articles": articles})
	}
}

// 判断是否管理员
// 判断用户是否登录
func isAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authInfo := getAuthInfo(ctx)
		if authInfo.Identify == "admin" {
			ctx.Next()
			return
		}
		articles := []Article{}
		articles = append(articles, Article{Title: "没有访问权限", Model: Model{CreatedAt: time.Now()}})
		ctx.Abort()
		ctx.HTML(200, "root/403.html", gin.H{"authInfo": authInfo, "articles": articles})
	}
}
