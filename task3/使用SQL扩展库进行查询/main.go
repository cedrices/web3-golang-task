package main

import (
	"fmt"
	"strconv"

	"github.com/cedrices/web3-golang-task/task3/link"
)

type Employees struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"comment:员工姓名"`
	Department string  `gorm:"comment:所属部门"`
	Salary     float64 `gorm:"comment:员工薪资"`
}

func main() {
	db := link.InitMysql()
	db.AutoMigrate(&Employees{})
	emp := []Employees{{Name: "李四", Department: "技术部", Salary: 8000},
		{Name: "张三", Department: "技术部", Salary: 6000},
		{Name: "王五", Department: "技术部", Salary: 18000},
		{Name: "赵六", Department: "技术部", Salary: 30000}}
	db.Create(&emp)
	var employee []Employees
	db.Model(&Employees{Department: "技术部"}).Find(&employee)
	for _, e := range employee {
		fmt.Println("员工姓名:", e.Name, "所属部门:", e.Department, "薪资:", strconv.FormatFloat(e.Salary, 'g', -1, 64))
	}

	// 按薪资降序查询技术部员工
	empl := Employees{}
	db.Model(&Employees{}).Order("salary desc").Limit(1).Find(&empl)
	fmt.Println("技术部薪资最高的员工:", empl.Name, "薪资:", fmt.Sprintf("%.2f", empl.Salary))
}
