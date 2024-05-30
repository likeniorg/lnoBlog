package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 登录路由
func login(r *gin.Engine) {
	r.POST("/login", func(ctx *gin.Context) {
		type Login struct {
			User User
		}
		login := Login{}
		err := ctx.ShouldBind(&login)
		if err != nil {
			ctx.String(501, "解析客户端参数失败")
			return
		}
		cliPass := login.User.Password
		// 查找用户名是否存在
		result := db.Where("username = ?", login.User.Username).Find(&login.User)
		if result.Error != nil {
			ctx.String(501, "查找用户数据失败")
			return
		}

		// 不存在这个用户
		if result.RowsAffected != 1 {
			ctx.String(501, "账号或密码错误")
			return
		}

		// 判断用户输入的密码是否对应其Hash值
		err = bcrypt.CompareHashAndPassword([]byte(login.User.Password), []byte(cliPass))
		if err != nil {
			ctx.String(501, "账号或密码错误")
			return
		}

		uidstring := strconv.Itoa(int(login.User.ID))
		cookie, err := GenerateJWT(uidstring, login.User.Identify, login.User.Username)
		if err != nil {
			ctx.String(501, "服务器生成Cookie失败")
			return
		}

		ctx.SetCookie("authorization", cookie, 3600, "/", "", false, false)
		ctx.String(200, "登录成功")

	})

}
