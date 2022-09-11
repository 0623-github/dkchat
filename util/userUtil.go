package util

import (
	"dkchat/model"
	"dkchat/sql"
	"fmt"
	"strconv"
)
func DistributeUserAccount(user *model.User) {
	number := sql.SelectMaxAccount()
	num := "1"
	if number != "" {
		num = stringAddOne(number)
	}
	user.AccountNumber = num
}

func stringAddOne(number string) string {
	num, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}
	num += 1
	return strconv.Itoa(num)
}

