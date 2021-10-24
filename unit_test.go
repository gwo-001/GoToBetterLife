package main

import (
	orm "GoToBetterLife/dal/database"
	"GoToBetterLife/dal/models"
	"GoToBetterLife/util"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	util.HasLength(1)

	util.GetNowDate()
}

func TestSelect(t *testing.T) {
	user:=models.User{}
	user.Username="tom"
	user.Password="123456"
	num:=0
	orm.Db.Model(&user).Where("username=? AND password=?",user.Username,user.Password).Count(&num)
	fmt.Println(num)
}
func TestMd5(t *testing.T)  {
	md5Ctx:=md5.New()

	test:="123456"
	md5Ctx.Write([]byte(test))

	fmt.Println(hex.EncodeToString(md5Ctx.Sum(nil)))
}
