package auth

import (
	"Backed/database/dal"
	"fmt"
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
	expiryTime := time.Now().Add(10 * time.Minute)
	result := dal.PostgreSQL.Exec(`
		DELETE FROM users 
		WHERE is_active = false 
		AND created_at < $1`,
		expiryTime)
	if result.Error != nil {
		fmt.Errorf("清理未在指定时间进行激活处理的账号失败: %s", result.Error)
	} else if rows := result.RowsAffected; rows > 0 {
		log.Printf("清理了 %d 个未在指定时间进行激活处理的账号", rows)
	}
}
