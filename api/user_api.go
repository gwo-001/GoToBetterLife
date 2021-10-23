package api

import (
	"GoToBetterLife/dal/models"
	"GoToBetterLife/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Users 返回所有的用户
func Users(c *gin.Context) {
	var user models.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到相关用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    result,
	})
}

// UserSignUp 添加数据
func UserSignUp(c *gin.Context) {
	userExist := false
	var err error
	user := models.User{}
	err = c.BindJSON(&user)

	//ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.Zero,
			util.Message:    util.Fail,
			util.Data:       err,
		})
	}

	// 这里检测用户是否已经注册了
	userExist, err = user.CountUserName()
	if userExist {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.Zero,
			util.Message:    util.Fail,
			util.Data:       "User already exist",
		})
		return
	}

	// 这里插入新用户
	if _,err=user.Insert();err == nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.One,
			util.Message:    util.Success,
			util.Data:       "user signUp success",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		util.StatusCode: util.Zero,
		util.Message:    util.Fail,
		util.Data:       "user signUp failed",
	})
}

// Update 根据用户id来修改用户密码
func Update(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 60)
	user.Password = c.Request.FormValue("password")

	//这里根据用户id和新密码来修改用户密码
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "未找到相关用户",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改用户成功",
	})

}

// Destroy 根据id来删除用户
func Destroy(c *gin.Context) {
	var user models.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除用户成功",
	})

}
