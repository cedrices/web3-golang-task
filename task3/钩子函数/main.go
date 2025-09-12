package main

import (
	"github.com/cedrices/web3-golang-task/task3/link"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"comment:用户名;unique"`
	Age       uint   `gorm:"comment:年龄"`
	PostCount int    `gorm:"comment:文章数量;default:0"`
	Post      []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `gorm:"comment:标题"`
	UserID       uint      `gorm:"comment:用户ID;index"`
	CommentState string    `gorm:"comment:评论状态;default:'有评论'"`
	Comment      []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"comment:评论内容"`
	PostID  uint   `gorm:"comment:帖子ID;index"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&count)
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", count).Error
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if count == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_state", "无评论").Error
	}
	return
}

func main() {
	db := link.InitMysql()
	// db.AutoMigrate(&User{}, &Post{}, &Comment{})
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// users := User{Name: "李四", Age: 45, Post: []Post{
	// 	{Title: "测试000", Comment: []Comment{
	// 		{Content: "评论1111"},
	// 		{Content: "评论2222"},
	// 	}},
	// 	{Title: "测试888", Comment: []Comment{{Content: "谢谢分享"}}},
	// }}
	// db.Create(&users)

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var comment Comment
	comment.PostID = 3
	db.Delete(&comment, "post_id = ?", 3)
}
