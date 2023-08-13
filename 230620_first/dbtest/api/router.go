package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//登录注册
	r.GET("/login", Login_get)
	r.POST("/login", Login)
	r.GET("/register", Register_get)
	r.POST("/register", Register)
	//book
	r.GET("/index", Index_get)
	v1 := r.Group("/index")
	{
		v1.GET("/serchbook", SerchBook_get)
		v1.POST("/serchbook/dan", GetBookDan)
		v1.POST("/serchbook/duo", GetBookDuo)
		v1.GET("/borrowbook", BorrowBook_get)
		v1.POST("/borrowbook", BorrowBook)
		v1.GET("/returnbook", ReturnBook_get)
		v1.POST("/returnbook", ReturnBook)
		v1.GET("/history", History_get)
		v1.POST("/history", History)
	}
	r.GET("/admin", AdminLog_get)
	r.POST("/admin", AdminLog)
	v2 := r.Group("/admin")
	{
		v2.GET("/index", AdminIndex_get)
		v2.POST("/index", AdminIndex_get)
		// 路径为 "/admin/index/add-book"
		v2.POST("/index/add-book", AdminIndex_addbook)
		// 路径为 "/admin/index/delete-book"
		v2.POST("/index/delete-book", AdminIndex_delbooks)
		// 路径为 "/admin/index/add-user"
		v2.POST("/index/add-user", AdminIndex_addreader)
		// 路径为 "/admin/index/delete-user"
		v2.POST("/index/delete-user", AdminIndex_delreader)
	}
}
