package repositories

import (
	"BBS/model"
	"BBS/pkg/snowflake"
	"errors"
	"github.com/kataras/iris/v12"
	"log"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (u *userRepository) CreateUser(username string, userPwd string, Email string) (err error) {
	Db := model.DB()
	uid := snowflake.GenID()
	user := &model.User{
		UserID:   uid,
		Username: username,
		Password: userPwd,
		Email:    Email,
	}
	if err := Db.Table("user").Create(&user).Error; err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func (u *userRepository) SelectUser(userName string) (*model.User, error) {
	User := new(model.User)
	if userName == "" {
		return &model.User{}, errors.New("用户名不能为空")
	}
	model.DB().Where("username=?", userName).Table("user").Find(&User)
	return User, nil
}

func (u *userRepository) IsExistUser(userName string) error {
	user, err := u.SelectUser(userName)
	if user != nil && err == nil {
		return errors.New("该用户名已存在")
	}
	return nil
}

func (*userRepository) GetUserByUid(Uid int64) (user *model.User, err error) {
	db := model.DB()
	db.Where("user_id = ?", Uid).Table("user").Find(&user)
	return user, nil

}

func (*userRepository) GetCurrentUserID(ctx iris.Context) (int64, error) {
	UserID, err := ctx.PostValueInt64("UserID")
	if err != nil {
		log.Fatalf("用户id获取失败")
	}
	return UserID, nil
}
