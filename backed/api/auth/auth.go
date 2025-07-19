package auth

import (
	"Backed/database/dal"
	"Backed/database/model"
	"Backed/utils"
	"database/sql"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const (
	errorKey      = "error"
	internalError = "服务器内部错误，请稍后再试"
	invalidCreds  = "用户名或密码错误"
	registerError = "无法注册用户，请换个名称/邮箱或稍后再试"
	tokenError    = "无法生成令牌，请稍后再试"
)

func Login(c *gin.Context) {
	// 表单获取
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证输入长度
	if len(username) > 64 || len(password) > 128 {
		c.JSON(http.StatusBadRequest, gin.H{errorKey: "无效的输入"})
		return
	}

	var user model.User
	err := dal.PostgreSQL.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{errorKey: "用户不存在"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}

	// 检验密码
	var match bool
	match, err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: err.Error()})
		return
	}
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{errorKey: invalidCreds})
		return
	}

	// 生成令牌
	token, err := utils.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: tokenError})
		return
	}

	// 返回令牌
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	// 获取表单
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	// 生成用户ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}

	id := node.Generate().Int64()

	// 使用Argon2加密密码
	var hashPassword string
	hashPassword, err = utils.HashPasswordWithArgon2(password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
	}

	// 检验是否已存在同一用户名/邮箱
	var exists model.User
	err = dal.PostgreSQL.Where("username = ?", username).First(&exists).Error
	if err == nil { // 如果存在同一用户名
		c.JSON(http.StatusConflict, gin.H{errorKey: "该名称已被占用"})
		return
	}

	err = dal.PostgreSQL.Where("email = ?", sql.NullString{String: email, Valid: true}).First(&exists).Error
	if err == nil { // 如果存在同一邮箱
		c.JSON(http.StatusConflict, gin.H{errorKey: "该邮箱已被占用"})
		return
	}

	// 插入用户
	newUser := model.User{
		UserID:   id,
		Username: username,
		Password: hashPassword,
		Email:    email,
	}

	err = dal.PostgreSQL.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: registerError})
		return
	}

	// 生成令牌
	token, err := utils.GenerateToken(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: tokenError})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "注册成功",
	})
}

func VerifyAuth(c *gin.Context) {
	token := c.Query("token")

	var verificationToken model.VerificationToken
	result := dal.PostgreSQL.Where("token = ? AND expires_at > ?", token, time.Now()).Preload("User").First(&verificationToken)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"title":   "验证失败",
			"message": "无效或过期的验证链接",
		})
		return
	}

	dal.PostgreSQL.Model(&verificationToken.User).Update("is_active", true)

	dal.PostgreSQL.Delete(&verificationToken)

	c.JSON(http.StatusBadRequest, gin.H{
		"title":   "aaa",
		"message": "aaa",
	})
}
