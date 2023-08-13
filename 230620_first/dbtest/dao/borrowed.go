package dao

import (
	"fmt"
	"time"
)

type Borrow struct {
	BookId       int    `gorm:"column:book_id"`
	ReadeAccount string `gorm:"column:reader_account"`
	BorrowBook   string `gorm:"column:borrow_book"`
	BorrowTime   string `gorm:"column:borrow_time"`
	ReturnTime   string `gorm:"column:return_time"`
}

func (Borrow) TableName() string {
	return "borrow"
}

func InsertBorrow(useraccount string, books Books) {
	var br Borrow
	BTime := time.Now().Format("2006-01-02 15:04:05")
	br = Borrow{
		BookId:       books.Id,
		ReadeAccount: useraccount,
		BorrowBook:   books.BookName,
		BorrowTime:   BTime,
	}
	err := DB.Create(br).Error
	if err != nil {
		fmt.Println(err)
	}
}

func UpdaBorrow(books Books) {
	var br Borrow
	BTime := time.Now().Format("2006-01-02 15:04:05")
	DB.Model(&br).Where("borrow_book=? AND book_id=?", books.BookName, books.Id).Update("return_time", BTime)
}

func BorHistory(useraccount string) ([]Borrow, bool) {
	var date []Borrow
	err := DB.Where("reader_account=?", useraccount).Find(&date).Error
	if err != nil {
		fmt.Println(err)
		return date, false
	}
	return date, true
}
