package middleware

import (
	"errors"
	"fmt"
	"time"

	"Tiktok/dao"
	"Tiktok/global"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	User_Id   int64
	User_Name string
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}
func (j *JWT) CreateClaims(baseClaims BaseClaims) *CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime,
			Issuer:    global.GVA_CONFIG.JWT.Issuer,
		},
	}
	return &claims
}

//生成token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //生成头|体|签名
	return token.SignedString(j.SigningKey)                    //进行加密
}

//解析token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

//通过用户的id和name生成token
func GetJWTByUserInfo(user dao.User) (token string) {
	basicClaims := BaseClaims{
		User_Id:   int64(user.ID),
		User_Name: user.Name,
	}
	claims := NewJWT().CreateClaims(basicClaims)
	token, err := NewJWT().CreateToken(*claims)
	if err != nil {
		//到时候换成日志
		fmt.Println("获取token失败", err.Error())
	}
	return token
}

//解析token获取带有id和name的baseClaims
func GetParseToken(token string) (baseClaims *BaseClaims, err error) {
	var claims *CustomClaims
	claims, err = NewJWT().ParseToken(token)
	if err != nil {
		return baseClaims, errors.New("解析token错误")
	}
	baseClaims = &(claims.BaseClaims)
	return
}
