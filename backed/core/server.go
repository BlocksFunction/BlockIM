package core

import (
	"Backed/api"
	"Backed/api/articles"
	"Backed/api/auth"
	"Backed/config"
	"Backed/database/dal"
	"Backed/utils"
	"Backed/utils/accout"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func StartServer() {
	// 加载配置文件
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("无法加载配置文件!")
	}

	// 加载MySQL数据库
	if err := dal.InitMySQL(cfg.Database); err != nil {
		log.Fatalf("无法使用MySQL: %s", err)
		return
	}

	err = accout.SendVerificationEmail("3398817447@qq.com", accout.GenerateToken(1946204647656001536))
	if err != nil {
		log.Fatal(err)
	}

	// 加载清理未激活账号程序
	go accout.CleanupNotActiveUserTask()

	router := gin.Default()
	initRoutes(router)
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("启动服务器失败: %v\n", err)
		os.Exit(1)
	}
}

func initRoutes(router *gin.Engine) {
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)
	router.GET("/verify", auth.VerifyAuth)
	router.GET("/me", utils.AuthMiddleware(), api.Me)

	articleGroup := router.Group("/article")
	articleGroup.GET("/list", articles.GetList)
	articleGroup.POST("/add", utils.AuthMiddleware(), articles.AddArticle)
	articleGroup.POST("/delete/:id", utils.AuthMiddleware(), articles.DeleteArticle)
	articleGroup.GET("/get/:id", articles.GetArticle)
}
