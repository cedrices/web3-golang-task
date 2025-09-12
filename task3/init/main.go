package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `gorm:"type:varchar(100)"`
	Age        int
	UserExt    UserExt `gorm:"embedded;embeddedPrefix:test_"`
	CreditCard CreditCard
}

type UserExt struct {
	Email   string `gorm:"column:unique_email;type:varchar(100)"`
	Address string `gorm:"size:500"`
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type Email struct {
	gorm.Model
	UserID uint
	Email  string `gorm:"uniqueIndex:idx_user_email"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.Age += 1
	fmt.Println("BeforeSave钩子函数 用户年龄+1后为:", user.Age)
	return
}

func (user *User) AfterSave(tx *gorm.DB) (err error) {

	fmt.Println("AfterSave钩子函数 用户年龄+1后为:", user.Age)
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Age += 1
	fmt.Println("BeforeCreate钩子函数 用户年龄+1后为:", user.Age)
	return
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate钩子函数 用户年龄+1后为:", user.Age)
	return
}

func main() {
	dsn := "dev_haos:haos123456@tcp(127.0.0.1:3306)/go_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引))
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	//自动创建表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&CreditCard{})

	//==================创建表记录==================begin==================
	// user := User{Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}
	//创建表记录
	// result := db.Create(&user)
	//创建记录为指定字段赋值
	//result := db.Select("Name", "Age").Create(&user)
	//创建记录忽略指定字段
	//result := db.Omit("Name").Create(&user)

	//批量创建记录
	// users := []User{{Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}, {Name: "李四", Age: 21, UserExt: UserExt{Email: "222@163.com", Address: "上海长虹区"}}, {Name: "王五", Age: 38, UserExt: UserExt{Email: "333@163.com", Address: "广东深圳"}}}
	// result := db.Create(&users)
	//一个insert sql语句插入语句上限 例如 CreateInBatches(&users, 2) 表示每个insert语句插入2条记录
	// result := db.CreateInBatches(&users, 5)

	//跳过钩子执行
	// result := db.Debug().Session(&gorm.Session{SkipHooks: true}).CreateInBatches(&users, 1)
	// for i, _ := range users {
	// 	users[i].ID = 0
	// }

	//result1 := db.Debug().Session(&gorm.Session{SkipHooks: true}).CreateInBatches(&users, 1)

	//返回插入的错误
	// fmt.Println(result.Error)
	//返回插入的记录数
	// fmt.Println(result.RowsAffected)
	//==================创建表记录==================end=================

	//=================查询表记录==================begin==================
	// var user1 User
	// var users []User

	//var user1 = User{Model: gorm.Model{ID: 15}}

	//查询第一条记录
	// db.Debug().First(&user1)
	// fmt.Println("查询到的第一条记录:", user1)
	// db.Debug().First(&user1, 2) //主键查询
	// fmt.Println("主键查询到的记录:", user1)
	// db.Debug().Take(&user1) //随便查询一条记录
	// fmt.Println("随便查询到的一条记录:", user1)
	// db.Debug().Last(&user1) //查询最后一条记录
	// fmt.Println("查询到的最后一条记录:", user1)
	// db.Debug().First(&user1, "name = ?", "张三") //条件查询第一条记录
	// fmt.Println("条件查询到的第一条记录:", user1)
	// db.Debug().First(&user1, "15") //条件查询第一条记录
	// fmt.Println("条件查询第一条记录:", user1)
	// db.Debug().First(&user1) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", user1)
	// db.Debug().Model(User{Model: gorm.Model{ID: 15}}).First(&user1) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", user1)
	// db.Debug().Find(&users, []uint{15, 4, 8}) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)
	// db.Debug().Find(&user1, "id = ?", 15) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", user1)
	// db.Where("name = ?", "张三").First(&user1)
	// fmt.Println("条件查询所有记录:", user1)

	// db.Where("name <> ?", "张三").Find(&users) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Where("name IN ?", []string{"张三", "李四"}).Find(&users) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Where("name LIKE ?", "张%").Find(&users)
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Where("name = ? AND age >= ?", "张三", "10").Find(&users)
	// fmt.Println("条件查询所有记录:", users)

	// db.Where("updated_at < ?", time.Now()).Offset(1).Limit(2).Find(&users)
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Where("created_at BETWEEN ? AND ?", time.DateTime, time.Now()).Find(&users)
	// fmt.Println("条件查询所有记录:", users)
	// var user = User{Model: gorm.Model{ID: 4}}
	// db.Debug().Where("id = ?", 20).Find(&user)

	// db.Debug().Where(&User{Name: "张三", Age: 18}).Find(&user1)
	// fmt.Println("条件查询所有记录:", user1)

	// db.Debug().Where(map[string]interface{}{"Name": "张三", "Age": 18}).Find(&users)
	// fmt.Println("条件查询所有记录:", users)

	// db.Where([]int64{15, 4, 8}).Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user.CreatedAt.Local().Format(time.DateTime))
	// }

	// db.Debug().Where(&User{Name: "张三"}, "name", "Age").Find(&users) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)
	// db.Debug().Where(&User{Name: "张三"}, "name").Find(&users) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Find(&users, "name = ?", "张三") //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Find(&users, &User{Age: 18}) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Not(&User{Name: "张三"}, "name", "Age").Find(&users) //条件查询所有记录
	// fmt.Println("条件查询所有记录:", users)

	// db.Debug().Not(map[string]interface{}{"Name": "张三", "Age": 18}).Find(&users)
	// fmt.Println("条件查询所有记录:", users)

	// db.Where("name = '李四'").Or(User{Name: "张三", Age: 18}).Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	// db.Debug().Where("name = '王五'").Or(map[string]interface{}{"Name": "张三", "Age": 18}).Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	// db.Select("name", "age").Find(&users, []uint{15, 4, 8}) //条件查询所有记录
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	// db.Debug().Select([]string{"name"}).First(&users, 15) //条件查询所有记录
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	// rows, err := db.Debug().Table("users").Select("COALESCE(age,?)", 42).Order("age desc").Rows()
	// if err != nil {
	// 	fmt.Println("查询出错:", err)
	// 	return
	// }
	// for rows.Next() {
	// 	var age int
	// 	err = rows.Scan(&age)
	// 	if err != nil {
	// 		fmt.Println("查询出错:", err)
	// 		return
	// 	}
	// 	fmt.Println("查询到的年龄:", age)
	// }
	// db.Debug().Clauses(clause.OrderBy{
	// 	Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	// }).Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	// db.Debug().Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "张三").Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }
	// rows, err := db.Debug().Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Having("sum(age) > ?", 50).Rows()

	// if err != nil {
	// 	fmt.Println("查询出错:", err)
	// 	return
	// }
	// for rows.Next() {
	// 	var date string
	// 	var total int
	// 	err = rows.Scan(&date, &total)
	// 	if err != nil {
	// 		fmt.Println("查询出错:", err)
	// 		return
	// 	}
	// 	fmt.Println("查询到的记录:", date, total)
	// }

	// type Result struct {
	// 	Date  string
	// 	Total int
	// }
	// var results []Result

	// db.Debug().Table("users").Select("date(created_at) as date, sum(age) as total").Group("date(created_at)").Having("sum(age) > ?", 50).Scan(&results)
	// for _, result := range results {
	// 	fmt.Println("查询到的记录:", result.Date, result.Total)
	// }

	// db.Debug().Distinct("name", "age").Order("name, age desc").Find(&users)
	// for _, user := range users {
	// 	fmt.Println("查询到的记录:", user)
	// }

	//db.Select("users.id", "emails.email").Joins("JOIN emails ON emails.user_id = users.id").Where("AND emails.email = ?", "111@example.org").Find(&users)

	// db.Debug().Find(&users, "name = ?", "张三") //条件查询所有记录
	// for i, _ := range users {
	// 	users[i].Age += 100
	// }
	// fmt.Println("条件查询所有记录:", users)
	// db.Debug().Session(&gorm.Session{SkipHooks: true}).Save(&users)
	//db.Debug().Model(&User{}).Where("id = ?", 15).Update("name", "六子")

	//db.Debug().Model(&User{}).Where(&User{Model: gorm.Model{ID: 15}}).Updates(map[string]interface{}{"name": "hello", "age": 18})

	//db.Debug().Model(&User{}).Omit("name").Where(&User{Model: gorm.Model{ID: 15}}).Updates(map[string]interface{}{"name": "hello", "age": 500})

	//db.Debug().Model(&User{Model: gorm.Model{ID: 15}}).Select("Name").Updates(User{Name: "new_name", Age: 0})

	//db.Debug().Model(&User{Model: gorm.Model{ID: 15}}).Select("*").Omit("Age").Updates(User{Name: "jinzhu", Age: 0})

	//db.Exec("UPDATE users SET name = ? where id = ?", "jinzhu", 15)

	//db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")

	//默认跳过 hooks
	// db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).UpdateColumn("Age", gorm.Expr("Age - ?", 10))

	// db.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("Email", db.Model(&Email{}).Select("user_id").Where("emails.user_id = users.id"))

	//默认逻辑删除
	//db.Debug().Delete(&User{}, 15)

	//物理删除
	//db.Debug().Unscoped().Delete(&User{}, 15)

	// db.Debug().Where("name LIKE ?", "%jinzhu%").Limit(3).Delete(&User{})

	// db.Debug().Delete(&User{}, "name LIKE ? and age like ?", "%jinzhu%", "%28%")

	//db.Exec("DELETE FROM users")

	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	// user := User{Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}
	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	// user := User{ID: 28, Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}
	// db.Debug().Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "Name"}},
	// 	DoUpdates: clause.Assignments(map[string]interface{}{"name": "铁柱"}),
	// }).Create(&user)

	// user := User{ID: 28, Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}
	// db.Session(&gorm.Session{SkipHooks: true}).Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "id"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	// }).Create(&user)

	// 全量更新 默认主键是否重复 重复 全量更新
	// user := User{ID: 15, Name: "王五", Age: 500, UserExt: UserExt{Email: "222@163.com", Address: "上海长虹区"}}
	// db.Debug().Clauses(clause.OnConflict{
	// 	UpdateAll: true,
	// }).Create(&user)

	// db.Debug().Create(&User{
	// 	Name: "王五", Age: 500, UserExt: UserExt{Email: "222@163.com", Address: "上海长虹区"},
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })

	// user := User{Name: "天王", Age: 300, UserExt: UserExt{Email: "333@163.com", Address: "广东深圳"}}
	// db.Omit("CreditCard").Create(&user)

	// var age int64
	// db.Raw("SELECT SUM(age) FROM users WHERE id = ?", "15").Scan(&age)
	// fmt.Println("查询到的年龄总和:", age)
	// db.Exec("DROP TABLE credit_cards")

	// stmt := db.Session(&gorm.Session{DryRun: true}).First(&User{}, 1).Statement
	// fmt.Println(stmt.SQL.String(), stmt.Vars)

	// rows, err := db.Model(&User{}).Where("name = ?", "王五").Select("name, age").Rows()
	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var age int
	// 	err = rows.Scan(&name, &age)
	// 	if err != nil {
	// 		fmt.Println("查询出错:", err)
	// 		return
	// 	}
	// 	fmt.Println("查询到的记录:", name, age)
	// }

	// rows, err := db.Raw("select name, age from users where name = ?", "张三").Rows()
	// defer rows.Close()
	// if err != nil {
	// 	fmt.Println("查询出错:", err)
	// 	return
	// }
	// for rows.Next() {
	// 	var name string
	// 	var age int
	// 	err = rows.Scan(&name, &age)
	// 	if err != nil {
	// 		fmt.Println("查询出错:", err)
	// 		return
	// 	}
	// 	fmt.Println("查询到的记录:", name, age)
	// }

}
