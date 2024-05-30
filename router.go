package main

import (
	"github.com/gin-gonic/gin"
)

// 路由
func router(router *gin.Engine) {
	publicPage(router)
	userPage(router)
	dataHandler(router)
	register(router)
	login(router)
	adminPage(router)
	adminHandler(router)
}
