package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// Argon2 参数
const (
	argonTime    = 1         // 迭代次数
	argonMemory  = 64 * 1024 // 使用的内存量 (KB)
	argonThreads = 4         // 使用的线程数
	argonKeyLen  = 32        // 生成的密钥长度 (字节)
	argonSaltLen = 16        // 盐的长度 (字节)
)

// hashPasswordWithArgon2 使用 Argon2 算法哈希密码
func HashPasswordWithArgon2(password string) (string, error) {
	// 生成随机盐
	salt := make([]byte, argonSaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("生成盐失败: %w", err)
	}

	// 使用 Argon2id 变体进行哈希
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonTime,
		argonMemory,
		argonThreads,
		argonKeyLen,
	)

	// 编码为存储格式: $argon2id$v=19$m=...,t=...,p=...$<base64(salt)>$<base64(hash)>
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		argonMemory,
		argonTime,
		argonThreads,
		b64Salt,
		b64Hash,
	), nil
}

// verifyPassword 验证密码是否匹配哈希值
func VerifyPassword(password, encodedHash string) (bool, error) {
	// 解析存储的哈希字符串
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("无效的哈希格式")
	}

	// 提取参数
	var version int
	if _, err := fmt.Sscanf(parts[2], "v=%d", &version); err != nil {
		return false, fmt.Errorf("解析版本号失败: %w", err)
	}

	var memory, iterations, parallelism uint32
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism); err != nil {
		return false, fmt.Errorf("解析参数失败: %w", err)
	}

	// 解码盐和哈希
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("解码盐失败: %w", err)
	}

	originalHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("解码哈希失败: %w", err)
	}

	// 使用相同参数哈希输入的密码
	hashToCompare := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		uint8(parallelism),
		uint32(len(originalHash)),
	)

	// 使用恒定时间比较防止计时攻击
	return subtle.ConstantTimeCompare(originalHash, hashToCompare) == 1, nil
}
