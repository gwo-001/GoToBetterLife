package models

import (
	orm "GoToBetterLife/api/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

// Insert 添加用户
func (user *User) Insert() (id int64, err error) {
	result := orm.Db.Create(&user)
	id = user.ID
	if err = result.Error; err != nil {
		return id,err
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
