package dao

import (
	"log"
)

type Reader struct {
	Readername string `gorm:"column:readername"`
	Gender     string `gorm:"column:gender"`
	Account    string `gorm:"column:account"`
	Password   string `gorm:"column:passwd"`
}

func (Reader) TableName() string {
	return "reader"
}

func SaveUser(user *Reader) {
	//数据库操作，首先连接数据库
	err := DB.Create(user).Error
	if err != nil {
		log.Println("inster error", err)
	}
}

// 查询
func GetById(id int64) Reader {
	var user Reader
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("get user by id error", err)
	}
	return user
}

func GetAll() []Reader {
	var user []Reader
	err := DB.Find(&user).Error
	if err != nil {
		log.Println("get ALL user error", err)
	}
	return user
}

// 更新
func UpdateUser(id int64) {
	err := DB.Model(&Reader{}).Where("id=?", id).Update("Username", "lide").Error
	if err != nil {
		log.Println("update user error", err)
	}
}

// 删除
func DeleteUser(id int64) {
	err := DB.Model(&Reader{}).Where("id=?", id).Delete(&Reader{}).Error
	if err != nil {
		log.Println("delete user error", err)
	}
}

func GetForLogin(account, password string) bool {
	var user Reader
	res := DB.Where("account=? AND passwd=?", account, password).First(&user)
	if res.RowsAffected != 0 {
		return true
	}
	return false
}

func GetForReg(reader *Reader) bool {
	res := DB.Where("account=?", reader.Readername).First(&reader)
	if res.RowsAffected != 0 {
		return false
	} else {
		DB.Create(&reader)
		return true
	}
}

func AddReader(reader Reader) bool {
	res := DB.Where("account=?", reader.Account).First(&reader)
	if res.RowsAffected != 0 {
		return false
	} else {
		DB.Create(&reader)
		return true
	}
}

func DeleteReader(reader Reader) bool {
	err := DB.Where("account=?", reader.Account).Delete(&reader).Error
	if err != nil {
		return false
	}
	return true
}
