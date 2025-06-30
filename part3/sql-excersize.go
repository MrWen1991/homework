package part3

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"homework/gorms"
)

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段
id （主键，自增）、
name （学生姓名，字符串类型）、
age （学生年龄，整数类型）、
grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Students struct {
	ID    uint32
	Name  string
	Age   uint32
	Grade string
}

var db *gorm.DB

func SqlExcersizeRun1() {
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	/*err := db.Save(&Students{Name: "张三", Age: 20, Grade: "三年级"}).Error
	if err != nil {
		panic(errors.New("create Student failure!"))
	}*/

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	s1 := &Students{}
	db.Debug().Where("age > ?", 18).Find(&s1)
	s1Bytes, _ := json.Marshal(s1)
	fmt.Printf("s1:%v\n", string(s1Bytes))

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级")

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Debug().Model(&Students{}).Where("age < ?", 15).Delete(&Students{})
}

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）
和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
type Account struct {
	gorm.Model
	Balance float32
}
type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float32
}

func SqlExcersizeRun2() {
	/*accounts := []Account{
		{Balance: 200},
		{Balance: 300},
	}
	db.CreateInBatches(&accounts, 2)*/

	var deduct float32 = 100

	db.Transaction(func(tx *gorm.DB) error {
		accs := []Account{}
		err := db.Debug().Where("id in ?", []int{1, 2}).Order("ID asc").Find(&accs).Error
		if err != nil {
			panic(errors.New("query account1 failed"))
		}
		if len(accs) != 2 {
			panic(errors.New("query result length is not 2"))
		}
		if accs[0].Balance < 100 {
			panic(errors.New("account A balance is not enough"))
		}
		accs[0].Balance -= deduct
		accs[1].Balance += deduct
		err = db.Debug().Select("balance").Updates(&accs[0]).Error
		err11 := db.Debug().Select("balance").Updates(&accs[1]).Error
		if err != nil || err11 != nil {
			fmt.Printf("update result failed,error:%v \n ", err.Error())
			panic(errors.New("update result failed"))
		}

		statement := db.Save(&Transaction{
			FromAccountID: accs[0].Model.ID,
			ToAccountID:   accs[1].Model.ID,
			Amount:        deduct,
		}).Statement
		if statement.RowsAffected != 1 {
			panic(errors.New("save transaction failed"))
		}
		return nil
	})
}

func init() {
	db = gorms.DB
	db.AutoMigrate(&Students{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
}
