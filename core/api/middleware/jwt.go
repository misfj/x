package middleware

import (
	"coredx/config"
	"coredx/log"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const authorizationHeader = "Authorization"

var (
	TokenExpired     = errors.New("token已过期,请重新登录。")
	TokenNotValidYet = errors.New("token无效,请重新登录。")
	TokenMalformed   = errors.New("token不正确,请重新登录。")
	TokenInvalid     = errors.New("这不是一个token,请重新登录。")
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getJwtFromHeader(c)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusForbidden,
				"msg":     "",
				"message": "no token",
			})
			c.Set("jwt_error", err.Error())
			c.Abort()
			return
		}
		claim, err := ParseToken(token, config.GetJwt().Secret)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusForbidden,
				"msg":  TokenMalformed.Error(),
				"data": "",
			})
			c.Set("jwt_error", err.Error())
			c.Abort()
			return
		}
		if errors.Is(claim.Valid(), nil) {
			c.Set("AppName", claim.AppName)
			//todo 加一层数据库token验证
			c.Next()
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusForbidden,
				"msg":     "",
				"message": "no token",
			})
			//解决token问题,直接返回。
			c.Set("jwt_error", err.Error())
			c.Abort()
			return
			// return TokenMalformed
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusForbidden,
				"msg":  TokenExpired.Error,
				"data": "",
			})
			c.Set("jwt_error", err.Error())
			c.Abort()
			return

		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusForbidden,
				"message": TokenInvalid.Error(),
				"data":    "",
			})
			c.Set("jwt_error", err.Error())
			c.Abort()
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusForbidden,
				"msg":  TokenNotValidYet.Error(),
				"data": "",
			})
			c.Set("jwt_error", err.Error())
			c.Abort()
			return
		}
		//fmt.Println("无语")
	}
}

func getJwtFromHeader(c *gin.Context) (string, error) {
	aHeader := c.Request.Header.Get(authorizationHeader)
	if len(aHeader) == 0 {
		return "", fmt.Errorf("token is empty")
	}
	strs := strings.SplitN(aHeader, " ", 2)
	if len(strs) != 2 || strs[0] != "Bearer" {
		return "", fmt.Errorf("token 不符合规则")
	}
	return strs[1], nil
}

type CustomClaims struct {
	AppName string `json:"app_name"`
	jwt.RegisteredClaims
}

func BuildClaims(exp time.Time, appName string) *CustomClaims {
	return &CustomClaims{
		AppName: appName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.GetJwt().Iss,
		},
	}
}

// GenToken 生成jwt token
func GenToken(c *CustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}

// ParseToken 解析jwt token
func ParseToken(jwtStr, secretKey string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
