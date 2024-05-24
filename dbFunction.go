package main

import "github.com/gin-gonic/gin"

// 获取全部文章版块
func getCategory() []ArticleCategory {
	categories := []ArticleCategory{}
	db.Find(&categories)
	return categories
}

// 获取用户信息
func getAuthInfo(ctx *gin.Context) AuthInfo {
	authInfo, exist := ctx.Get("AuthInfo")
	if !exist {
		return AuthInfo{}
	}
	return authInfo.(AuthInfo)

}
