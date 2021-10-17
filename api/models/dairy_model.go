package models

import (
	orm "GoToBetterLife/api/database"
	"fmt"
	"time"
)

type Dairy struct {
	id             int64  `json:"id" gorm:"column:id;notnull;type:int primary key auto_increment;comment:'主键'"`
	Datetime       int    `json:"datetime" gorm:"column:date_time;type:int;comment:'日期'"`
	Article        string `json:"article" gorm:"column:article;type:varchar(255);comment:'日记本体'"`
	BeOnDutyTime   string `json:"be_on_duty_time" gorm:"column:be_on_duty_time;type:varchar(8);comment:'开始工作的时间'"`
	OffDutyTime    string `json:"off_duty_time" gorm:"column:off_duty_time;type:varchar(8);comment'下班时间'"`
	IsWorkOutToday string `json:"is_workout_today" gorm:"column:is_workout_today:type:varchar(8):comment:'是否运动了'"'`
}

var Dairies []Dairy

// CreateDairy 建立dairy表
func (dairy *Dairy) CreateDairy() (result int, err error) {
	database := orm.Db
	result = 0

	if database.HasTable(&Dairy{}) {
		fmt.Println("[CreateDairy] table dairy already exist")
		return
	}
	if err = database.AutoMigrate(Dairy{}).Error; err != nil {
		fmt.Println("[CreateDairy] creat table dairy error")
		return
	}
	result = 1
	return
}

// GetLatestTenDairies 获取到最近十天的日记，用于页面刚刚加载的时候返回给到前端
func (dairy *Dairy) GetLatestTenDairies() (dairies []Dairy, err error) {
	if err = orm.Db.Find(&dairies).Limit(10).Error; err != nil {
		return
	}
	return
}

//GetNewPage 这里需要分页返回每10天的记录
func (dairy *Dairy) GetNewPage(page int) (dairies []Dairy, err error) {
	// 这里分别获取到当前的年月日，用来存库
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()
	var nowData = year*10000 + month*100 + day
	// 这里分别获取到需要返回页的其实日期的数据和结束日期的数据
	var endDate = nowData - (page-1)*10
	var startDate = endDate - 10

	err = orm.Db.Where("datetime BETWEEN ? AND ?", endDate, startDate).Find(&dairies).Error
	if err != nil {
		return nil, err
	}
	return
}

//InsertNewDairies 用来记录今天的上班下班时间、是否锻炼身体、是否学习等
func (dairy *Dairy) InsertNewDairies(dairy1 Dairy) (dateTime int, err error) {
	result := orm.Db.Create(dairy1)
	err = result.Error
	if err != nil {
		return 0, err
	}
	dateTime = dairy.Datetime
	return
}

// DeleteDairies 删除掉一段时间的日记
func (dairy *Dairy) DeleteDairies(startTime int, endTime int) (deleteNum int, err error) {
	result := orm.Db.Where("date_time &gt;= ? and data_time &lt;= ?", startTime, endTime).Delete(dairy)
	err = result.Error
	if err != nil {
		fmt.Println("[DeleteDairies] failed")
		return 0, err
	}
	return 1, err
}
