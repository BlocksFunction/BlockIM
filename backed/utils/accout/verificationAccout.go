package accout

import (
	"Backed/config"
	"Backed/database/dal"
	"Backed/database/model"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"log"
	"time"
)

func CleanupNotActiveUserTask() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		cleanupInactiveUsers()
	}
}

func cleanupInactiveUsers() {
	expiryTime := time.Now().Add(-10 * time.Minute)
	result := dal.PostgreSQL.Exec(`
		DELETE FROM users 
		WHERE is_active = false 
		AND created_at < $1`,
		expiryTime)

	if result.Error != nil {
		log.Printf("清理未激活用户失败: %s", result.Error)
	} else {
		rows := result.RowsAffected
		if rows > 0 {
			fmt.Printf("清理了 %d 个未激活用户", rows)
		}
	}
}

func GenerateToken(userID int64) string {
	token := uuid.NewV4().String()

	verificationToken := model.VerificationToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	dal.PostgreSQL.Create(&verificationToken)

	return token
}

func SendVerificationEmail(email, token string) error {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("无法加载配置文件!")
	}
	log.Printf("SMTP配置: Host=%s, Port=%d, User=%s", cfg.SMTP.SMTPHost, cfg.SMTP.SMTPPort, cfg.SMTP.SMTPUser)

	verifyURL := fmt.Sprintf("%s/verify?id=%s", cfg.App.AppHost, token)
	body := fmt.Sprintf("<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"UTF-8\">\n    <title>激活您的%s账号</title>\n    <style>\n        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; line-height: 1.6; }\n        .container { max-width: 600px; margin: 20px auto; padding: 20px; border: 1px solid #e0e0e0; border-radius: 8px; }\n        .header { background-color: #f8f9fa; padding: 20px; text-align: center; border-radius: 8px 8px 0 0; }\n        .content { padding: 30px 20px; }\n        .button { display: inline-block; padding: 12px 24px; background-color: #007bff; color: white !important; \n                 text-decoration: none; border-radius: 4px; font-weight: bold; margin: 20px 0; }\n        .footer { text-align: center; padding: 20px; color: #6c757d; font-size: 0.9em; }\n        .note { background-color: #fff3cd; padding: 15px; border-radius: 4px; margin: 20px 0; }\n    </style>\n</head>\n<body>\n    <div class=\"container\">\n        <div class=\"header\">\n            <h2>欢迎使用 %s！</h2>\n        </div>\n        \n        <div class=\"content\">\n            <p>你好！</p>\n            <p>感谢您注册 %s 账号 <strong>%s</strong>。请点击下方按钮激活您的账号：</p>\n            \n            <p style=\"text-align: center;\">\n                <a href=\"%s\" class=\"button\">激活账号</a>\n            </p>\n            \n            <div class=\"note\">\n                <p><strong>重要提示：</strong></p>\n                <p>此激活链接将在 <strong style=\"color: #d9534f;\">10分钟</strong> 后过期</p>\n                <p>如果按钮无法点击，请复制以下链接到浏览器：</p>\n                <p><a href=\"%s\">%s</a></p>\n            </div>\n            \n            <p>如果你未注册%s，请忽略此邮件。</p>\n        </div>\n        \n        <div class=\"footer\">\n            <p>&copy; %s. 保留所有权利。</p>\n            <p><a href=\"%s\">访问网站</a>\n        </div>\n    </div>\n</body>\n</html>", cfg.App.Name, cfg.App.Name, cfg.App.Name, email, verifyURL, verifyURL, verifyURL, cfg.App.Name, cfg.App.Name, cfg.App.FrontHost)

	m := gomail.NewMessage()
	m.SetHeader("From", "imaur@foxmail.com")
	m.SetHeader("To", "3398817447@qq.com")
	m.SetHeader("Subject", "请验证你的账号")
	m.SetBody("text/html", body)
	d := gomail.NewDialer(cfg.SMTP.SMTPHost, cfg.SMTP.SMTPPort, cfg.SMTP.SMTPUser, cfg.SMTP.SMTPPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
