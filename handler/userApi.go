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
	Result bool `json:"result"`
}

// SignUp func
func SignUp(ctx *gin.Context) {
	req := &signUpRequest{}
	resp := signUpResponse{
		Code: util.StatusSuccess,
		Message: "success",
		Result: true,
	}
	fmt.Println(ctx)
	err := ctx.BindJSON(req)
	if err != nil {
		fmt.Println(err)
		resp.Code = util.StatusParamErr
		resp.Message = "注册参数错误"
		resp.Result = false
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
	Result bool `json:"result"`
	Token string `json:"token"`
}

func SignIn(ctx *gin.Context) {
	req := &signInRequest{}
	resp := &signInResponse{
		Code: util.StatusSuccess,
		Message: "登陆成功",
		Result: true,
	}
	req.AccountNumber = ctx.Query("account_number")
	req.Pwd = ctx.Query("pwd")
	fmt.Println(req)
	user := &model.User{
		AccountNumber: req.AccountNumber,
	}
	sql.SelectUserByAccount(user)
	fmt.Println(user.Pwd, req.Pwd)
	if req.Pwd == "" {
		resp.Code = util.StatusPwdErr
		resp.Message = "密码为空"
		resp.Result = false
	} else if user.Pwd != req.Pwd {
		resp.Code = util.StatusPwdErr
		resp.Message = "密码错误，请重试"
		resp.Result = false
	} else {
		resp.NickName = user.NickName
		resp.Grade = user.Grade
		resp.AccountNumber = user.AccountNumber
		resp.Result = true
		token, err := util.GeneralToken(util.UserToUserClaims(user))
		resp.Token = token
		if err != nil {
			resp.Code = util.StatusGeneralTokenErr
			resp.Message = "后端生成token错误，请重试"
			resp.Result = false
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