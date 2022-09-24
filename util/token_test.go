package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)
var userClaims = &UserClaims{
	AccountNumber: "123",
	Pwd: "123",
	RealName: "123",
	NickName: "123",
	Age: 1,
	Grade: 1,
	Gender: 1,
	Phone: "123",
	StandardClaims: jwt.StandardClaims{},
}
func TestGeneralToken(t *testing.T) {
	token, err := GeneralToken(userClaims)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token, _ := GeneralToken(userClaims)
	parseToken := ParseToken(token)
	fmt.Println(parseToken)
}