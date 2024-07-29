package middware

import (
	"coredx/log"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const (
	expire_msg    = "token is expired"
	token_illegal = "that's not even a token"
)

func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		claims, err := ParseToken(token)
		if err != nil {
			response.Response(http.StatusForbidden, err.Error(), "", ctx)
			ctx.Abort()
			return
		}
		log.Debugf("claims:%v", claims)
		//用claims的用户名和密码去数据库进行比对
		//直接通过

		if verify(claims.UserName, claims.Issuer, claims.Subject) {
			if claims.UserName == "internal" {
				ctx.Next()
			} else {
				//todo 查询数据库
				response.Response(http.StatusForbidden, "not allow", "", ctx)
				ctx.Abort()
				return
			}

		} else {
			response.Response(http.StatusForbidden, "verify failed", "", ctx)
			ctx.Abort()
			return
		}
	}
}

type MyClaims struct {
	UserName             string
	jwt.RegisteredClaims // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

// var MySecret []byte
//
//	func Secret() jwt.Keyfunc {
//		return func(token *jwt.Token) (interface{}, error) {
//			return MySecret, nil // 这是我的secret
//		}
//	}
func GenerateToken(exp int, iss string, sub string, userName string) (tokenString string, err error) {
	claim := MyClaims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp) * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                        // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                        // 生效时间
			Issuer:    iss,
			Subject:   sub,
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(MySecret)
	return tokenString, err
}

func ParseToken(tokenss string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &MyClaims{}, Secret())
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(token_illegal)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(expire_msg)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
func verify(username, iss, sub string) bool {
	//if model.QueryByUsername(mysql.NwlyDB, username) == nil {
	//	return iss == config.JwtConfig().Iss && sub == config.JwtConfig().Sub
	//} else {
	//	return false
	//}
	return iss == config.JwtConfig().Iss && sub == config.JwtConfig().Sub
}
