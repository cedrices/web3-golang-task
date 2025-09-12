package main

import (
	"fmt"
	"strconv"

	"github.com/cedrices/web3-golang-task/task3/link"
)

type Book struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"comment:书名"`
	Author string  `gorm:"comment:作者"`
	Price  float64 `gorm:"comment:价格"`
}

func main() {
	db := link.InitMysql()
	db.AutoMigrate(&Book{})
	db.Create(&Book{Title: "Go语言编程", Author: "张三", Price: 59.99})
	var books []Book
	db.Where("price > ?", 50).Find(&books)
	for _, book := range books {
		fmt.Println("书名:", book.Title, "作者:", book.Author, "价格:", strconv.FormatFloat(book.Price, 'g', -1, 64))
	}
}
