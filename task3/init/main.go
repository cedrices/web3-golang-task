package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)"`
	Age     int
	UserExt UserExt `gorm:"embedded;embeddedPrefix:test_"`
}

type UserExt struct {
	Email   string `gorm:"column:unique_email;type:varchar(100)"`
	Address string `gorm:"size:500"`
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

	//==================创建表记录==================begin==================
	//user := User{Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}
	//创建表记录
	//result := db.Create(&user)
	//创建记录为指定字段赋值
	//result := db.Select("Name", "Age").Create(&user)
	//创建记录忽略指定字段
	//result := db.Omit("Name").Create(&user)

	//批量创建记录
	//users := []User{{Name: "张三", Age: 18, UserExt: UserExt{Email: "111@163.com", Address: "北京市朝阳区"}}, {Name: "李四", Age: 21, UserExt: UserExt{Email: "222@163.com", Address: "上海长虹区"}}, {Name: "王五", Age: 38, UserExt: UserExt{Email: "333@163.com", Address: "广东深圳"}}}
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
	// fmt.Println(result.Error, result1.Error)
	//返回插入的记录数
	// fmt.Println(result.RowsAffected, result1.RowsAffected)
	//==================创建表记录==================end=================

	//=================查询表记录==================begin==================
	//var user1 User
	var users []User

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

}
