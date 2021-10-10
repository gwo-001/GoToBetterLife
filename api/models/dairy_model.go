package models

import (
	orm "GoToBetterLife/api/database"
	"time"
)

type Dairy struct {
	Datetime       int  `json:"datetime"`
	Article        string `json:"article"`
	BeOnDutyTime   string `json:"be_on_duty_time"`
	OffDutyTime    string `json:"off_duty_time"`
	IsWorkOutToday string `json:"is_workout_today"`
}

var Dairies []Dairy

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
func (dairy *Dairy) InsertNewDairies(datetime int64,)(result string) {
	result := orm.Db.Create()

}
