package main

import (
	"fmt"

	"github.com/cedrices/web3-golang-task/task3/link"
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	Name  string
	Age   uint
	Grade string
}

func main() {
	db := link.InitMysql()

	db.AutoMigrate(&Students{})
	db.Create(&Students{Name: "张三", Age: 20, Grade: "一年级"})
	var student []Students
	db.Where(" age > ?", 18).Find(&student)
	for _, stu := range student {
		fmt.Println("学生姓名:", stu.Name, "年龄:", stu.Age, "年级:", stu.Grade)
	}
	db.Model(&Students{}).Where("name = ?", "张三").Update("Grade", "四年级")
	db.Where("age < ?", 15).Delete(&Students{})
}
