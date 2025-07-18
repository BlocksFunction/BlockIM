package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

const (
	authHeader   = "Authorization"
	bearerPrefix = "Bearer "
	tokenCtxKey  = "jwtClaims"
)

var (
	jwtKey          = []byte("secret")
	errNotLoggedIn  = gin.H{"error": "你还没有登录，无法进行此操作"}
	errInvalidToken = gin.H{"error": "无效的访问令牌"}
)

func GenerateToken(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaderVal := c.GetHeader(authHeader)
		if authHeaderVal == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errNotLoggedIn)
			return
		}

		tokenStr := strings.TrimPrefix(authHeaderVal, bearerPrefix)
		if tokenStr == authHeaderVal { // 如果没有Bearer前缀
			c.AbortWithStatusJSON(http.StatusUnauthorized, errInvalidToken)
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&jwt.RegisteredClaims{},
			func(*jwt.Token) (interface{}, error) { return jwtKey, nil },
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errInvalidToken)
			return
		}

		// 存储解析后的Claims
		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
			c.Set(tokenCtxKey, claims)
		}

		c.Next()
	}
}

// GetClaims 从Context获取Claims
func GetClaims(c *gin.Context) *jwt.RegisteredClaims {
	if claims, exists := c.Get(tokenCtxKey); exists {
		return claims.(*jwt.RegisteredClaims)
	}
	return nil
}
