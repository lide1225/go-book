package api

import (
	"dbt/dao"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SaveUser(c *gin.Context) {
	user := &dao.Reader{
		// Username:   "zhangsan",
		// Password:   "123456",
	}
	dao.SaveUser(user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {

	// user := dao.GetById(1)
	user := dao.GetAll()
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateUser(2)
	user := dao.GetById(2)
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	dao.DeleteUser(2)
	user := dao.GetById(2)
	c.JSON(200, user)
}

// ===================Login=====================
func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if dao.GetForLogin(account, password) {
		//session获取用户信息
		session := sessions.Default(c)      // 获取session
		session.Set("useraccount", account) // 设置session的值
		session.Save()                      // 保存session
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号或密码错误",
		})
	}
}

func Login_get(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//===============================================

// ===================register===================
func Register(c *gin.Context) {
	// name := c.PostForm("name")
	gender := c.PostForm("gender")
	readername := c.PostForm("readername")
	account := c.PostForm("account")
	password := c.PostForm("password")
	reader := &dao.Reader{
		Readername: readername,
		Gender:     gender,
		Account:    account,
		Password:   password,
	}

	if dao.GetForReg(reader) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号已经存在",
		})
	}
}

func Register_get(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
