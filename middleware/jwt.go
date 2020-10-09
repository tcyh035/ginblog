package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var code int

// JwtKey jwtkey切片
var JwtKey = []byte(utils.JwtKey)

// MyClaims 用户Claims
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10)
	setClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	requestClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := requestClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.Error
	}

	return token, errmsg.Success
}

// VerifyToken 验证Token
func VerifyToken(token string) (*MyClaims, int) {
	var myClaims MyClaims
	setToken, _ := jwt.ParseWithClaims(token, &myClaims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if claims, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return claims, errmsg.Success
	}

	return nil, errmsg.Error
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ErrorTokenNotExist
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong

			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage(code),
			})
			c.Abort()
			return
		}

		key, tCode := VerifyToken(checkToken[1])
		if tCode == errmsg.Error {
			code = errmsg.ErrorTokenTypeWrong

			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage(code),
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime

			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrorMessage(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
