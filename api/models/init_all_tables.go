package models

// InitAllTables 初始化所有数据表
func InitAllTables() (err error) {
	var user User
	var dairy Dairy
	num := 0

	if _, err = user.CreatUserTable(); err != nil {
		num += 1
	}
	if _, err = dairy.CreateDairy(); err != nil {
		num += 1
	}

	return  err
}
