package util

import (
	"dkchat/model"
	"fmt"
	"testing"
)

func TestStringAddOne(t *testing.T) {
	fmt.Println(stringAddOne("123"))
}

func TestDistributeUserAccount(t *testing.T) {
	user := &model.User{}
	DistributeUserAccount(user)
	fmt.Println(user.AccountNumber)
}
