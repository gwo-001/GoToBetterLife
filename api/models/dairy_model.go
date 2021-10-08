package models

import orm "GoToBetterLife/api/database"

type Dairy struct {
	Datetime       string `json:"datetime"`
	Article        string `json:"article"`
	BeOnDutyTime   string `json:"be_on_duty_time"`
	OffDutyTime    string `json:"off_duty_time"`
	IsWorkOutToday string `json:"is_work_out_today"`
}

var Dairies []Dairy

// GetTenDairies 获取到最近十天的日记
func (dairy *Dairy) GetTenDairies() (dairies []Dairy,err error){
	if err=orm.Db.Find(&dairies).Limit(10).Error;err!= nil {
		return
	}
	return
}

func (dairy *Dairy) AddNewDairies(article string,isworkouttoday string)  {

}
