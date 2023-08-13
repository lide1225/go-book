package api

import (
	"dbt/dao"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index_get(c *gin.Context) {
	session := sessions.Default(c)
	useraccount := session.Get("useraccount")
	c.HTML(200, "index.html", gin.H{
		"useraccount": useraccount,
	})
}

func SerchBook_get(c *gin.Context) {
	c.HTML(200, "serchbook.html", nil)
}

var books []dao.Books

func GetBookDan(c *gin.Context) {
	// 从表单中获取查询条件
	title := c.PostForm("title")
	searchtype := c.PostForm("selectBox")
	// 根据查询条件进行数据库的查询
	books, a := dao.GetBookByT(searchtype, title)
	// 返回查询结果
	if !a {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "查询失败",
		})
	} else {
		c.HTML(http.StatusOK, "serchbook.html", gin.H{
			"Books": books,
		})
	}
}

func GetBookDuo(c *gin.Context) {
	// 从表单中获取查询条件
	bookname := c.PostForm("bookname[]")
	author := c.PostForm("author[]")
	isbn := c.PostForm("isbn[]")
	subword := c.PostForm("subword[]")
	pubhouse := c.PostForm("pubhouse[]")
	// 根据查询条件进行数据库的查询
	DuoBooks, a := dao.GetBookByD(bookname, author, isbn, subword, pubhouse)
	// 返回查询结果
	if !a {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "查询失败",
		})
	} else {
		c.HTML(http.StatusOK, "serchbook.html", gin.H{
			"Books": DuoBooks,
		})
	}
}

// =================borrow==========================
func BorrowBook_get(c *gin.Context) {
	books, x := dao.HotBook()
	if x {
		c.HTML(200, "borrow.html", gin.H{
			"books": books,
		})
	}
}

func BorrowBook(c *gin.Context) {
	bookname := c.PostForm("bookname")
	author := c.PostForm("author")
	session := sessions.Default(c)
	useraccount := session.Get("useraccount").(string)
	if dao.BorBook(useraccount, bookname, author) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "借阅成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "借阅失败,此书正在被借阅或无此书",
		})
	}

}

//=====================return==========================

func ReturnBook_get(c *gin.Context) {
	c.HTML(200, "return.html", nil)
}

func ReturnBook(c *gin.Context) {
	bookname := c.PostForm("bookname")
	author := c.PostForm("author")
	if dao.RetBook(bookname, author) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "归还成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "归还失败,此书没有被借阅或无此书",
		})
	}
}

//======================history======================

func History_get(c *gin.Context) {
	c.HTML(200, "history.html", nil)
}

func History(c *gin.Context) {
	session := sessions.Default(c)
	useraccount := session.Get("useraccount").(string)
	history, a := dao.BorHistory(useraccount)
	if a {
		c.HTML(200, "history.html", gin.H{
			"history": history,
		})
	}

}
