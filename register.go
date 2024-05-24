package main

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// 注册账户路由
func register(router *gin.Engine) {
	router.POST("/register", func(ctx *gin.Context) {
		type Register struct {
			User User
			Key  string `validate:"required"`
		}

		// 接收客户端信息
		register := Register{}
		err := ctx.ShouldBind(&register)
		if err != nil {
			ctx.String(501, "信息解析失败")
			return
		}

		// 验证客户端信息
		validate := validator.New()
		err = validate.Struct(register)
		if err != nil {
			ctx.String(501, "填写信息不符合要求")
			return
		}

		// 判断注册密钥：临时策略，后期修改为一次性邀请密钥
		if register.Key != Key {
			ctx.String(501, "注册密钥错误")
			return
		}

		users := User{}
		isAdmin := db.Find(&users)
		if isAdmin.RowsAffected == 0 {
			register.User.Identify = "admin"
		} else {
			register.User.Identify = "user"

		}
		// 创建数据表
		register.User.CreatedAt = time.Now()
		hashPass, err := bcrypt.GenerateFromPassword([]byte(register.User.Password), bcrypt.MinCost)
		if err != nil {
			ctx.String(501, "密码格式错误，不要加入特殊字符")
			return
		}
		register.User.Password = string(hashPass)
		result := db.Create(&register.User)

		if result.Error != nil {
			if strings.Contains(result.Error.Error(), "Duplicate entry") {
				ctx.String(501, "已存在的用户名，请重新输入")
				return
			}

			ctx.String(501, result.Error.Error())
			return
		}

		sendMail(register.User.Email, []byte("注册物联网知识在线问答社区成功，欢迎访问！"))
		ctx.String(200, "注册成功")
	})
}
