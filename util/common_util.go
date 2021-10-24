package util

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"time"
)

// HasLength 用来判断输入i是否为nil或者为0值
func HasLength(i interface{}) (result bool) {
	result = false
	if i == nil {
		return result

	}
	iType := reflect.TypeOf(i).String()
	switch iType {
	case "int":
		iInt := i.(int)
		if iInt == 0 {
			result = true
		}
		return result
	case "string":
		iString := i.(string)
		if iString == "" {
			return
		}
		result = true
		return
	default:
		return
	}
	return
}

// GetNowDate 获取到当前时间20211020
func GetNowDate() (dateTime int) {
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()
	dateTime = year*10000 + month*100 + day
	return dateTime
}

// MD5 给用md5算法加密
func MD5(metaPassword string) (encryptPassword string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(metaPassword))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
