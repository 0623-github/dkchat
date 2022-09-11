package handler

import (
	"dkchat/model"
	"dkchat/sql"
	"dkchat/util"
	"fmt"
	"github.com/gin-gonic/gin"
)
// signUpRequest var
type signUpRequest struct {
	Pwd string `json:"pwd"`
	RealName string `json:"real_name"`
	NickName string `json:"nick_name"`
	Age int `json:"age"`
	Gender int `json:"gender"`
	Phone string `json:"phone"`
}
// signUpResponse var
type signUpResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	AccountNumber string `json:"account_number"`
}

// SignUp func
func SignUp(ctx *gin.Context) {
	req := &signUpRequest{}
	resp := signUpResponse{
		Code: util.StatusSuccess,
		Message: "success",
	}
	fmt.Println(ctx)
	err := ctx.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		resp.Code = util.StatusParamErr
		resp.Message = "注册参数错误"
	} else {
		u := signUpReqToUser(req)
		util.DistributeUserAccount(u)
		sql.NewUser(u)
		if u.AccountNumber != "" {
			resp.AccountNumber = u.AccountNumber
		}
	}
	ctx.JSON(200,&resp)
}

// signInRequest var
type signInRequest struct {
	AccountNumber string `json:"account_number"`
	Pwd string `json:"pwd"`
}
// signInResponse var
type signInResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	AccountNumber string `json:"account_number"`
	NickName string `json:"nick_name"`
	Grade int `json:"grade"`
}

func SignIn(ctx *gin.Context) {
	req := &signInRequest{}
	resp := &signInResponse{
		Code: util.StatusSuccess,
		Message: "登陆成功",
	}
	err := ctx.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		resp.Code = util.StatusParamErr
		resp.Message = "注册参数错误"
	} else {
		user := &model.User{
			AccountNumber: req.AccountNumber,
		}
		sql.SelectUserByAccount(user)
		fmt.Println(user.Pwd, req.Pwd)
		if req.Pwd == "" {
			resp.Code = util.StatusPwdErr
			resp.Message = "密码为空"
		} else if user.Pwd != req.Pwd {
			resp.Code = util.StatusPwdErr
			resp.Message = "密码错误，请重试"
		} else {
			resp.NickName = user.NickName
			resp.Grade = user.Grade
			resp.AccountNumber = user.AccountNumber
		}
	}
	ctx.JSON(200, resp)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func signUpReqToUser(req *signUpRequest) *model.User {
	user := model.User{
		Pwd: req.Pwd,
		RealName: req.RealName,
		NickName: req.NickName,
		Age: req.Age,
		Gender: req.Gender,
		Phone: req.Phone,
	}
	return &user
}