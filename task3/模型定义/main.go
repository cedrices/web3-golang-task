package main

import "github.com/cedrices/web3-golang-task/task3/link"

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"comment:用户名;unique"`
	Age  uint   `gorm:"comment:年龄"`
	Post []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID      uint      `gorm:"primaryKey"`
	Title   string    `gorm:"comment:标题"`
	UserID  uint      `gorm:"comment:用户ID;index"`
	Comment []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"comment:评论内容"`
	PostID  uint   `gorm:"comment:帖子ID;index"`
}

func main() {
	db := link.InitMysql()
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
