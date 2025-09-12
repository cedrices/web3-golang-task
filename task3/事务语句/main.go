package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Account struct {
	ID      uint    `gorm:"primaryKey;comment:账户ID主键"`
	Balance float64 `gorm:"comment:账户余额"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountID uint    `gorm:"comment:转出账户ID"`
	ToAccountID   uint    `gorm:"comment:转入账户ID"`
	Amount        float64 `gorm:"comment:转账金额"`
}

func main() {
	dsn := "dev_haos:haos123456@tcp(127.0.0.1:3306)/go_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("数据库连接成功", db)
	db.AutoMigrate(&Account{}, &Transaction{})
	// db.Create(&Account{Balance: 10000})
	// db.Create(&Account{Balance: 200})
	db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&fromAccount, 1).Error; err != nil {
			return err
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&toAccount, 2).Error; err != nil {
			return err
		}
		amount := 100.0
		if fromAccount.Balance < amount {
			return fmt.Errorf("余额不足")
		}
		fromAccount.Balance -= amount
		toAccount.Balance += amount
		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}
		transaction := Transaction{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        amount,
		}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
}
