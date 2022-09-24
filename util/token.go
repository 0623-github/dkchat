package util

import (
	"dkchat/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	AccountNumber string `json:"account_number"`
	Pwd string `json:"pwd"`
	RealName string `json:"real_name"`
	NickName string `json:"nick_name"`
	Age int `json:"age"`
	Grade int `json:"grade"`
	Gender int `json:"gender"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}
var(
	secret = []byte("16849841325189456f487")
	effectTime = 2 * time.Hour
)

func UserToUserClaims(user *model.User) *UserClaims {
	return &UserClaims{
		AccountNumber: user.AccountNumber,
		Pwd: user.Pwd,
		RealName: user.RealName,
		NickName: user.NickName,
		Age: user.Age,
		Grade: user.Grade,
		Gender: user.Gender,
		Phone: user.Phone,
		StandardClaims: jwt.StandardClaims{},
	}
}

func GeneralToken(claims *UserClaims) (string, error) {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, err
}

func VerifyToken(token *string) bool {
	if token != nil && len(*token) != 0 {
		return true
	}
	return false
}

func ParseToken(token string) *UserClaims {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := tokenClaims.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}