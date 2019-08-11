package controller

import (
	"../model"
	"../service"
	"../util"
	"fmt"
	"math/rand"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	pass := request.PostForm.Get("pass")

	user, err := userService.Login(mobile, pass)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}

}

var userService service.UserService

func Register(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	pass := request.PostForm.Get("pass")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_MAN
	user, err := userService.Register(mobile, pass, nickname, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
