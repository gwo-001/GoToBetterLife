package api

import (
	"GoToBetterLife/dal/models"
	"GoToBetterLife/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// LatestDairies 返回所有的日记
func LatestDairies(c *gin.Context) {
	var dairy models.Dairy
	var err error
	dairy.Datetime, err = strconv.ParseInt(c.Request.FormValue("datetime"), 10, 64)
	dairy.Article = c.Request.FormValue("article")
	result, err := dairy.GetLatestTenDairies()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.One,
			util.Message:    util.Fail,
			util.Data:       err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		util.StatusCode: util.One,
		util.Fail:       util.Fail,
		util.Data:       result,
	})
}

// AddNewDairy 添加一篇新的日记
func AddNewDairy(c *gin.Context) {
	var dairy models.Dairy
	var err error

	//dairy.Datetime, err = strconv.ParseInt(c.Request.FormValue("dateTime"), 10, 64)
	//dairy.Article = c.Request.FormValue("article")
	//dairy.BeOnDutyTime = c.Request.FormValue("beOnDutyTime")
	//dairy.OffDutyTime = c.Request.FormValue("offDutyTime")
	//dairy.IsWorkoutToday = c.Request.FormValue("isWorkoutToday")

	err = c.BindJSON(&dairy)
	ioutil.ReadAll(c.Request.Body)
	if !util.HasLength(dairy.Datetime) {
		dairy.Datetime = int64(util.GetNowDate())
	}
	dateTime, err := dairy.InsertNewDairies()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.One,
			util.Message:    util.Fail,
			util.Data:       err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		util.StatusCode: util.One,
		util.Message:    util.Success,
		util.Data:       dateTime,
	})
}
