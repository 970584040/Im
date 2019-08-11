package controller

import (
	"../args"
	"../model"
	"../service"
	"../util"
	"net/http"
)

var contactService service.ContactService

func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	_ = util.Bind(req, &arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w, users, len(users))
}

func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	_ = util.Bind(req, &arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w, comunitys, len(comunitys))
}
func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg

	//如果这个用的上,那么可以直接
	_ = util.Bind(req, &arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)

	//刷新用户群组
	AddGroupId(arg.Userid, arg.Dstid)

	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "")
	}
}

func Createcommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community

	_ = util.Bind(req, &arg)
	comm, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, comm, "ok")
	}
}

//
func Addfriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	_ = util.Bind(req, &arg)
	//调用service
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	//
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "好友添加成功")
	}
}
