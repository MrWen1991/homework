package part3

import (
	"fmt"
	"homework/gorms"
)

/*
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
//var dbx = gorms.DBX

type Employees struct {
	ID         int32   `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func (e Employees) Value() []interface{} {
	return []interface{}{e.Name, e.Department, e.Salary}
}

func SqlxRun1() {
	// 初始化数据
	/*employees := []Employees{
		{Name: "zhangsan", Department: "技术部", Salary: 16000},
		{Name: "lisi", Department: "技术部", Salary: 17000},
		{Name: "wangwu", Department: "技术部", Salary: 18000},
	}
	gorms.DB.Save(&employees)*/

	// TODO sqlx.In() 有问题
	/*tuples := convertUsersToTuples(employees)
	sql, args, _ := sqlx.In("insert into employees (id,name,department,salary) values (?)", tuples[0], tuples[1], tuples[2])
	_, err := gorms.DBX.Exec(sql, args...)
	if err != nil {
		panic(err)
	}*/
	//db.Exec()

	//使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	queryMultiRes := []Employees{}
	gorms.DBX.Select(&queryMultiRes, "select * from employees")
	fmt.Printf("query all employees info,length:%v \n", len(queryMultiRes))

	//使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	querySingleRes := Employees{}
	gorms.DBX.Get(&querySingleRes, "select * from employees order by salary desc limit 1")
	fmt.Printf("query highest salary employee info,salary:%v \n", querySingleRes.Salary)
}

/*
题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float32
}

func SqlxRun2() {
	// 初始化数据
	books := []Book{
		{Title: "jane aye", Author: "jane austine", Price: 32.5},
		{Title: "基督山伯爵", Author: "大仲马", Price: 52.5},
		{Title: "红楼梦", Author: "曹雪芹", Price: 85},
	}
	gorms.DB.Save(&books)

	queryMulBook := []Book{}
	gorms.DBX.Select(&queryMulBook, "select * from books where price > ?", 50)
	fmt.Printf("query books which price > 50, result:%v \n", queryMulBook)
}

func init() {
	gorms.DB.AutoMigrate(&Employees{})
	gorms.DB.AutoMigrate(&Book{})
	//employeeSql := "create table employees(\n    id int primary key,\n    name varchar(64) ,\n    department varchar(64),\n    salary decimal\n);"
	//exec, err := gorms.DBX.Exec(employeeSql)
	//if err != nil {
	//	panic(err)
	//}
	//affected, err := exec.RowsAffected()
	//fmt.Printf("dbx init affect rows %v \n", affected)
}
