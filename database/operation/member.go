package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetUserByUserName(userName string) (model.Member, error) {
	user := model.Member{}
	if err := db.Table(constant.TableMember).Where("username=?", userName).First(&user).Error; !util.DBQueryErr(err) {
		log.Printf("[GetPasswordByUserName] First err,userName:%+v,err:%+v", userName, err)
		return user, err
	}
	return user, nil
}

func CreateUser(userName string, password string, name string, identity string, tel string, email string, state int) error {
	user := model.Member{
		Username: userName,
		Password: password,
		Name:     name,
		Identity: identity,
		Tel:      tel,
		Email:    email,
		State:    state,
		Balance:  0.0,
	}

	if err := db.Table(constant.TableMember).Create(&user).Error; err != nil {
		log.Printf("[CreateUser] Create err,user:%+v,err:%+v", user, err)
		return err
	}

	return nil
}
