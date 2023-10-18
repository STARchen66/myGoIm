package models

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"myGoIm/utils"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name       string
	PassWord   string
	Phone      string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email      string `valid:"email"`
	Avatar     string //头像
	Identity   string
	ClientIp   string
	ClientPort string
	Salt       string

	LoginTime     sql.NullTime
	HeartbeatTime sql.NullTime
	LoginOutTime  sql.NullTime `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	//通过调用Unix()方法，我们可以从返回的时间对象中提取一个整数值，即当前时间的Unix时间戳。
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user

}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}

// 查找某个用户
func FindByID(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}
