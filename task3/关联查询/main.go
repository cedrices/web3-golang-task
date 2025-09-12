package main

import (
	"fmt"

	"github.com/cedrices/web3-golang-task/task3/link"
)

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
	// users := User{Name: "王五", Age: 45, Post: []Post{
	// 	{Title: "测试222", Comment: []Comment{
	// 		{Content: "评论1111"},
	// 		{Content: "评论2222"},
	// 		{Content: "评论3333"},
	// 	}},
	// 	{Title: "测试333", Comment: []Comment{{Content: "谢谢分享"}}},
	// }}
	// db.Create(&users)
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Post").Preload("Post.Comment").Find(&user, 1)
	fmt.Println("用户ID：", user.ID)
	fmt.Println("用户姓名：", user.Name)
	fmt.Println("用户年龄：", user.Age)
	for _, post := range user.Post {
		fmt.Println("  帖子ID：", post.ID)
		fmt.Println("  帖子标题：", post.Title)
		for _, comment := range post.Comment {
			fmt.Println("    评论ID：", comment.ID)
			fmt.Println("    评论内容：", comment.Content)
		}
	}
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post Post
	db.Model(&Post{}).Where(" id = (?)",
		db.Model(&Comment{}).Select("post_id").Group("post_id").Order("count(post_id) desc").Limit(1),
	).First(&post)
	fmt.Println("评论数量最多的文章ID：", post.ID)
	fmt.Println("评论数量最多的文章标题：", post.Title)

}
