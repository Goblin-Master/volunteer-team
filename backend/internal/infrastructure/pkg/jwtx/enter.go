package jwtx

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Role int

const (
	COMMON_USER = iota + 1
	INTERNAL_USER
	ADMIN
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	Role   Role  `json:"role"`
	jwt.RegisteredClaims
}

type Claims struct {
	UserID int64 `json:"user_id"`
	Role   Role  `json:"role"`
}

var (
	ErrDefault          = errors.New("jwt default error")
	ErrTokenEmpty       = errors.New("token is empty")
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrPermissionDenied = errors.New("permission denied")
)

func GenToken(c Claims) (string, error) {
	secret := configs.Conf.Auth.AccessSecret
	expiredTime := configs.Conf.Auth.AccessExpire
	// 创建一个我们自己的声明
	claims := MyClaims{
		c.UserID,
		c.Role,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiredTime) * time.Second)), // 过期时间
			Issuer:    "Volunteer-Team",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(secret))
}

func ParseToken(c *gin.Context) (int64, Role, error) {
	data := c.GetHeader("Authorization")
	if data == "" {
		return 0, 0, ErrTokenEmpty
	}
	token := strings.TrimPrefix(data, "Bearer ")
	// 解析token
	var claims MyClaims
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Conf.Auth.AccessSecret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return 0, 0, ErrTokenExpired
		}
		if strings.Contains(err.Error(), "signature is invalid") {
			return 0, 0, ErrTokenInvalid
		}
		if strings.Contains(err.Error(), "token contains an invalid") {
			return 0, 0, ErrTokenInvalid
		}
		fmt.Println(err)
		return 0, 0, ErrDefault
	}
	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims.UserID, claims.Role, nil
	}
	return 0, 0, ErrDefault
}

// 必须使用了鉴权中间件才能用
func GetUserID(c *gin.Context) int64 {
	if data, exists := c.Get(global.TOKEN_USER_ID); exists {
		user_id, ok := data.(int64)
		if ok {
			return user_id
		}
	}
	return 0
}

// 必须使用了鉴权中间件才能用
func GetRole(c *gin.Context) Role {
	if data, exists := c.Get(global.TOKEN_ROLE); exists {
		role, ok := data.(Role)
		if ok {
			return role
		}
	}
	return 0
}
