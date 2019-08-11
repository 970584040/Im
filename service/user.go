package service

import (
	"../model"
	"../util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct{}

//注册
func (s *UserService) Register(mobile, password, nickname, avatar, sex string) (user model.User, err error) {
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}

	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已注册！")
	}

	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(password, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngin.InsertOne(&tmp)

	return tmp, err
}

func (s *UserService) Login(mobile, password string) (user model.User, err error) {
	tmp := model.User{}
	_, _ = DbEngin.Where("mobile = ?", mobile).Get(&tmp)

	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在!")
	}

	if util.ValidatePasswd(password, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}

	str := fmt.Sprintf("%d", time.Now().Unix())
	tmp.Token = util.MD5Encode(str)
	_, err = DbEngin.ID(tmp.Id).Cols("token").Update(&tmp)
	if err != nil {
		tmp.Token = ""
		return tmp, errors.New("登陆失败")
	}
	fmt.Println(tmp, 123456)
	fmt.Printf("%s", tmp)
	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(
	userId int64) (user model.User) {

	//首先通过手机号查询用户
	tmp := model.User{}
	_, _ = DbEngin.ID(userId).Get(&tmp)
	return tmp
}
