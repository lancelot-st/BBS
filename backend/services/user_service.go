package services

import (
	"BBS/backend/repositories"
	"BBS/model"
	"BBS/pkg/encrypt"
	"log"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (u *userService) Register(userName string, userPwd string, Email string) bool {
	uPassword, errEncrypt := encrypt.Encrypt([]byte(userPwd), encrypt.Key)
	if errEncrypt != nil {
		log.Fatal(errEncrypt)
		return false
	}
	err := repositories.UserRepository.CreateUser(userName, uPassword, Email)
	if err != nil {
		return false
	}

	return true
}

func (u *userService) IsPwdSuccess(userName string, Pwd string) bool {
	user, err := repositories.UserRepository.SelectUser(userName)
	if err != nil {
		log.Println("用户名错误，没有该用户")
		return false
	}
	//解密
	pwd, err := encrypt.Decrypt(user.Password, encrypt.Key)
	if err != nil {
		log.Fatal("解密失败")
	}
	if pwd != Pwd {
		log.Fatal("用户密码不对")
		return false
	}

	return true
}

func (u *userService) GetUserByUidServ(Uid int64) (*model.User, error) {
	return repositories.UserRepository.GetUserByUid(Uid)
}

func (u *userService) GetUserByName(userName string) (*model.User, error) {
	user, err := repositories.UserRepository.SelectUser(userName)
	if err != nil {
		log.Fatal("用户获取失败")
	}
	return user, nil
}
