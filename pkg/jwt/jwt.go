package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
)

var mySingleKey = []byte("xiongzhao.com")

//claims结构体，也就是Payload中的claims

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userID int64, userName string) (string, error) {
	//创建要求
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := CustomClaims{
		UserID:   userID,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间，必须设置
			Issuer:    "XZ",                                  // 可不必设置，也可以填充用户名，
		},
	}
	jwtToken.Claims = claims
	//使用指定签名签名对象
	return jwtToken.SignedString(mySingleKey)
}

//解析token

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySingleKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func OnError(ctx iris.Context, err error) {
	if err == nil {
		return
	}

	ctx.StopExecution()

	ctx.JSON(err.Error())
}

func JWTAuthMiddleware() *jwt2.Middleware {
	return jwt2.New(jwt2.Config{
		ValidationKeyGetter: func(token *jwt2.Token) (interface{}, error) {
			return mySingleKey, nil
		},
		ErrorHandler:  OnError,
		SigningMethod: jwt2.SigningMethodHS256,
	})

}
