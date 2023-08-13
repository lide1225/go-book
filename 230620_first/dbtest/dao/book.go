package dao

import "fmt"

type Books struct {
	Id       int    `gorm:"column:id"`
	Bornum   int    `gorm:"column:bornum"`
	Borfre   int    `gorm:"column:borfre"`
	BookName string `gorm:"column:bookname"`
	Author   string `gorm:"column:author"`
	ISBN     string `gorm:"column:isbn"`
	Subword  string `gorm:"column:subword"`
	PubHouse string `gorm:"column:pubhouse"`
}

func (Books) TableName() string {
	return "books"
}

// 根据传来类型查询书
func GetBookByT(title, x string) ([]Books, bool) {
	var books []Books
	err := DB.Where(title+" LIKE ?", "%"+x+"%").Find(&books).Error
	if err != nil {
		return books, false
	} else if len(books) == 0 {
		return books, false
	}
	return books, true
}

func GetBookByD(bookname, author, isbn, subword, pubhouse string) ([]Books, bool) {
	var books []Books
	tx := DB.Model(&books)
	if bookname != "" {
		tx = tx.Where("bookname = ?", bookname)
	}
	if author != "" {
		tx = tx.Where("author = ?", author)
	}
	if isbn != "" {
		tx = tx.Where("isbn = ?", isbn)
	}
	if subword != "" {
		tx = tx.Where("subword = ?", subword)
	}
	if pubhouse != "" {
		tx = tx.Where("pubhouse = ?", pubhouse)
	}
	err := tx.Find(&books).Error
	if err != nil {
		return books, false
	} else if len(books) == 0 {
		return books, false
	}
	return books, true
}

func BorBook(useraccount, bookname, author string) bool {
	var books Books
	res := DB.Where("bookname=? AND author=?", bookname, author).First(&books)
	if res.RowsAffected != 0 && books.Bornum == 1 { //查询到书籍后要确定此书是否正在被借阅
		books.Borfre += 1
		DB.Model(&books).Where("bookname=? AND author=?", bookname, author).Update("bornum", 0)
		DB.Model(&books).Where("bookname=? AND author=?", bookname, author).Update("borfre", books.Borfre)
		InsertBorrow(useraccount, books) //操作borrow表
		return true
	}
	return false
}

func RetBook(bookname, author string) bool {
	var books Books
	res := DB.Where("bookname=? AND author=?", bookname, author).First(&books)
	if res.RowsAffected != 0 && books.Bornum == 0 { //查询到书籍后要确定此书是否正在被借阅
		books.Borfre += 1
		DB.Model(&books).Where("bookname=? AND author=?", bookname, author).Update("bornum", 1)
		UpdaBorrow(books)
		return true
	}
	return false
}

func AddBooks(books Books) bool {
	books.Bornum = 1
	err := DB.Create(&books).Error
	if err != nil {
		return false
	}
	return true
}

func DeleteBooks(books Books) bool {
	err := DB.Where("bookname=? AND author=?", books.BookName, books.Author).Delete(&books).Error
	if err != nil {
		return false
	}
	return true
}

func HotBook() ([]Books, bool) {
	var books []Books
	err := DB.Order("borfre DESC").Limit(10).Find(&books).Error
	if err != nil {
		fmt.Println(err)
		return books, false
	}
	return books, true
}
