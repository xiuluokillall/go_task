package task3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 使用gorm连接数据库
func GormConn() (db *gorm.DB) {
	dsn := "ben:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256, //string类型默认长度
		//DisableDatetimePrecision:  true,  //禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  //重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  //用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, //根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

// 使用sqlx连接数据库
func SqlxConn() (db *sqlx.DB, err error) {
	dsn := "ben:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, innerErr := sqlx.Connect("mysql", dsn)
	if innerErr != nil {
		return nil, innerErr
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	return db, nil
}

type Student struct {
	Id    int
	Name  string
	Age   int
	Grade string
}

func Q1() {
	db := GormConn()
	db.AutoMigrate(&Student{})
	student := &Student{
		Name:  "李四",
		Age:   14,
		Grade: "1年级",
	}
	db.Create(student)
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func Q2() {
	db := GormConn()
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	query := &Student{}
	db.Where("age >= ?", 18).Find(query)
	fmt.Println(query)
}

// Q3 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func Q3() {
	db := GormConn()
	query := &Student{}
	db.Model(query).Where("name = ?", "张三").Update("grade", "四年纪")
	fmt.Println(query)
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func Q4() {
	db := GormConn()
	query := &Student{}
	result := db.Where("age < ?", 15).Delete(query).Row()
	fmt.Println(query)
	fmt.Println(result)
}

/*
事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
type Account struct {
	Id      int
	Balance float64
}

type Transaction struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func Q5() {
	db := GormConn()
	//db.AutoMigrate(&Account{})
	//db.AutoMigrate(&Transaction{})
	//
	//accounts := &[]Account{
	//	{
	//		Balance: 50,
	//	},
	//	{
	//		Balance: 100,
	//	},
	//}
	//db.Create(accounts)

	err := db.Transaction(func(tx *gorm.DB) error {
		r1 := &Account{}
		r2 := &Account{}
		if err := tx.Where("id = ?", 1).First(r1).Error; err != nil {
			return err
		}
		aBalance := r1.Balance
		fmt.Println(aBalance)
		//if aBalance < 100 {
		//	panic("A账户余额不足")
		//}
		aBalance = float64(aBalance - 100)
		if err := tx.Where("id = ?", 2).First(r2).Error; err != nil {
			return err
		}
		bBalance := r2.Balance
		bBalance = float64(bBalance + 100)

		fmt.Printf("aBalance is %f", aBalance)
		fmt.Printf("bBalance is %f", bBalance)
		if err := tx.Model(&Account{}).Where("id = ?", 2).Update("balance", bBalance).Error; err != nil {
			return err
		}
		if err := tx.Model(&Account{}).Where("id = ?", 1).Update("balance", aBalance).Error; err != nil {
			return err
		}

		t := &Transaction{
			FromAccountId: r1.Id,
			ToAccountId:   r2.Id,
			Amount:        100,
		}
		tx.Create(t)
		return nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

type Employee struct {
	Id         int
	Name       string
	Department string
	Salary     float64
}

func Q6() {

}
