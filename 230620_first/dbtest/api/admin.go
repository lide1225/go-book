package api

import (
	"dbt/dao"

	"github.com/gin-gonic/gin"
)

func AdminLog_get(r *gin.Context) {
	r.HTML(200, "admin.html", nil)
}

func AdminLog(r *gin.Context) {
	account := r.PostForm("account")
	password := r.PostForm("password")
	if dao.AdminLogin(account, password) {
		r.JSON(200, gin.H{
			"msg": "登录成功",
		})
	} else {
		r.JSON(400, gin.H{
			"msg": "登录失败",
		})
	}
}

func AdminIndex_get(r *gin.Context) {
	r.HTML(200, "adminindex.html", nil)
}

func AdminIndex_addbook(r *gin.Context) {
	var books dao.Books
	books = dao.Books{
		BookName: r.PostForm("bookname"),
		Author:   r.PostForm("author"),
		ISBN:     r.PostForm("isbn"),
		Subword:  r.PostForm("subword"),
	}
	if dao.AddBooks(books) {
		r.JSON(200, gin.H{
			"msg": "添加图书成功",
		})
	} else {
		r.JSON(200, gin.H{
			"msg": "添加图书失败",
		})
	}
}

func AdminIndex_delbooks(r *gin.Context) {
	var books dao.Books
	books = dao.Books{
		BookName: r.PostForm("bookname"),
		Author:   r.PostForm("author"),
		ISBN:     r.PostForm("isbn"),
		Subword:  r.PostForm("subword"),
	}
	if dao.DeleteBooks(books) {
		r.JSON(200, gin.H{
			"msg": "删除图书成功",
		})
	} else {
		r.JSON(200, gin.H{
			"msg": "删除图书失败",
		})
	}
}

func AdminIndex_addreader(r *gin.Context) {
	var reader dao.Reader
	reader = dao.Reader{
		Readername: r.PostForm("username"),
		Gender:     r.PostForm("gender"),
		Account:    r.PostForm("account"),
		Password:   r.PostForm("password"),
	}
	if dao.AddReader(reader) {
		r.JSON(200, gin.H{
			"msg": "添加读者成功",
		})
	} else {
		r.JSON(200, gin.H{
			"msg": "添加读者失败",
		})
	}
}

func AdminIndex_delreader(r *gin.Context) {
	var reader dao.Reader
	reader = dao.Reader{
		Readername: r.PostForm("username"),
		Gender:     r.PostForm("gender"),
		Account:    r.PostForm("account"),
		Password:   r.PostForm("password"),
	}
	if dao.DeleteReader(reader) {
		r.JSON(200, gin.H{
			"msg": "删除读者成功",
		})
	} else {
		r.JSON(200, gin.H{
			"msg": "删除读者失败",
		})
	}
}
