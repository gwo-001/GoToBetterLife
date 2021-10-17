package models

import (
	orm "GoToBetterLife/api/database"
	"fmt"
)

type User struct {
	ID       int64  `json:"id" gorm:"column:id;notnull;type:int primary key auto_increment;comment:'主键'"`
	Username string `json:"username" gorm:"column:username; type:varchar(30); comment:'用户名'"`
	Password string `json:"password " gorm:"column:password; type:varchar(50); comment:'用户密码'"`
}

var Users []User

// CreatUserTable 建立user表
func (user *User) CreatUserTable() (result int, err error) {
	result = 0
	dataBase := orm.Db
	if dataBase.HasTable(&User{}) {
		fmt.Println("[CreatUserTable] 'user' table already exist")
		return result, err
	}
	err = dataBase.AutoMigrate(User{}).Error
	if err != nil {
		fmt.Println("[CreatUserTable] creat 'user' table error")
		return result, err
	}
	result = 1
	return result, err
}

// Insert 添加用户
func (user *User) Insert() (id int64, err error) {
	result := orm.Db.Create(&user)
	id = user.ID
	if err = result.Error; err != nil {
		return id, err
	}
	return
}

// Users 返回所有的user列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return
}

// Update 根据id来修改user
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.Db.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2：修改用的数据
	if err = orm.Db.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}

//Destroy 根据id来删除某用户
func (user *User) Destroy(id int64) (Result User, err error) {
	if err = orm.Db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Db.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
