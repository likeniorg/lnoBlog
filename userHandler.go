package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 用户功能逻辑处理
func userHandler() {
	user := r.Group("/user", isUser())
	{
		user.POST("/update_user", func(ctx *gin.Context) {
			user := User{}
			err := ctx.ShouldBind(&user)
			if err != nil {
				ctx.String(501, "参数解析失败")
				return
			}
			authInfo := getAuthInfo(ctx)
			uID, _ := strconv.Atoi(authInfo.Uid)
			user.ID = uID
			user.Identify = authInfo.Identify
			result := db.Model(&user).Update("username", user.Username)

			if result.Error != nil {
				ctx.String(501, "修改名字失败")
				return
			}

			cookie, err := GenerateJWT(authInfo.Uid, user.Identify, user.Username)
			if err != nil {
				ctx.String(501, "服务器生成Cookie失败")
				return
			}
			ctx.SetCookie("authorization", cookie, 3600, "/", "", false, false)
			ctx.String(200, "修改名字成功")
		})

		user.POST("/update_email", func(ctx *gin.Context) {
			user := User{}
			err := ctx.ShouldBind(&user)
			if err != nil {
				ctx.String(501, "参数解析失败")
				return
			}
			authInfo := getAuthInfo(ctx)
			uID, _ := strconv.Atoi(authInfo.Uid)
			user.ID = uID
			user.Identify = authInfo.Identify
			result := db.Model(&user).Update("email", user.Email)

			if result.Error != nil {
				ctx.String(501, "修改邮箱失败")
				return
			}
			ctx.String(200, "修改邮箱成功")
		})

		user.POST("/update_pass", func(ctx *gin.Context) {
			type Pass struct {
				Oldpass string
				Newpass string
			}

			pass := Pass{}
			err := ctx.ShouldBind(&pass)
			if err != nil {
				ctx.String(501, "参数解析失败")
				return
			}
			user := User{}
			authInfo := getAuthInfo(ctx)
			uID, _ := strconv.Atoi(authInfo.Uid)

			db.Where("id = ?", uID).Find(&user)
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass.Oldpass))
			if err != nil {
				ctx.String(501, "旧密码错误"+err.Error())
				return
			}
			newpass, _ := bcrypt.GenerateFromPassword([]byte(pass.Newpass), bcrypt.DefaultCost)
			result := db.Model(&user).Update("password", newpass)

			if result.Error != nil {
				ctx.String(501, "修改密码失败")
				return
			}
			ctx.String(200, "修改密码成功")
		})
	}

}
