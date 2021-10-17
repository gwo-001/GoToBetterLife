package api

import (
	"GoToBetterLife/dal/models"
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
		"data":result,
	})
}

// Store 添加数据
func Store(c *gin.Context) {
	var user models.User
	user.ID, _ = strconv.ParseInt(c.Request.FormValue("id"),10,64)
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加用户成功",
		"id":      id,
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
