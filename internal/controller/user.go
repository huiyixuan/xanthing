package controller

import (
	"xanthing/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserCtl struct{}

// GetUserInfo 获取用户信息示例，演示参数处理工具类的使用
func (ctl *UserCtl) GetUserInfo(c *gin.Context) {
	// 使用参数处理工具类
	params, err := utils.NewRequestParams(c)
	if err != nil {
		utils.BadRequest(c, "参数解析失败: "+err.Error())
		return
	}

	// 获取各种类型的参数
	userId := params.GetString("user_id")
	username := params.GetString("username")
	age := params.GetInt("age")
	score := params.GetFloat("score")
	isActive := params.GetBool("is_active")

	// 检查必要参数是否存在
	if userId == "" {
		utils.BadRequest(c, "user_id参数不能为空")
		return
	}

	// 获取所有参数（用于调试或日志记录）
	allParams := params.GetAll()

	// 返回处理结果
	utils.Success(c, gin.H{
		"message":    "用户信息获取成功",
		"user_id":    userId,
		"username":   username,
		"age":        age,
		"score":      score,
		"is_active":  isActive,
		"all_params": allParams,
	})
}

// CreateUser 创建用户示例，演示POST请求参数处理
func (ctl *UserCtl) CreateUser(c *gin.Context) {
	params, err := utils.NewRequestParams(c)
	if err != nil {
		utils.BadRequest(c, "参数解析失败: "+err.Error())
		return
	}

	// 获取用户信息参数
	username := params.GetString("username")
	email := params.GetString("email")
	password := params.GetString("password")
	role := params.GetString("role")

	// 验证必要参数
	if username == "" || email == "" || password == "" {
		utils.BadRequest(c, "用户名、邮箱和密码不能为空")
		return
	}

	// 模拟创建用户逻辑
	userInfo := gin.H{
		"id":       1,
		"username": username,
		"email":    email,
		"role":     role,
		"status":   "active",
	}

	utils.SuccessWithMsg(c, "用户创建成功", userInfo)
}
