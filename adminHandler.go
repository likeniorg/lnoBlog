package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 处理管理员权限功能
func adminHandler(r *gin.Engine) {
	admin := r.Group("/admin", isAdmin())
	{
		// 删除用户
		admin.POST("/del_user", func(ctx *gin.Context) {
			user := User{}
			err := ctx.ShouldBind(&user)
			if err != nil {
				ctx.String(501, "解析参数失败")
				return
			}
			result := db.Delete(&user)
			if result.Error != nil {
				ctx.String(501, "删除用户失败")
				return
			}
			ctx.String(200, "删除用户成功")
		})

		// 管理员发送邮件给用户
		admin.POST("/send_email", func(ctx *gin.Context) {
			// 收件人信息
			type EmailInfo struct {
				ID        int       `gorm:"column:id;primarykey"`
				Recipient string    `validate:"required,email"`
				Title     string    `validate:"required"`
				Content   string    `validate:"required"`
				CreatedAt time.Time `gorm:"column:created_at"`
			}
			info := EmailInfo{}
			if err := ctx.ShouldBind(&info); err != nil {
				ctx.String(401, "客户端信息解析失败")
				return
			}
			// 验证信息是否符合格式
			validate := validator.New()
			if err := validate.Struct(info); err != nil {
				ctx.String(501, "填写信息不符合要求")
				return
			}

			// result := db.Create(&info)
			// if result.Error != nil {
			// 	ctx.String(501, "记录邮件失败")
			// 	return
			// }

			// 发送邮件
			err := sendMail(info.Recipient, info.Title, info.Content)
			if err != nil {
				ctx.String(501, "发送邮件失败")
				return
			}

			ctx.String(200, "发送邮件成功")
		})
	}

}
