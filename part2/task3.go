package main

import "fmt"

func main() {
	/*	var s Shape = &Rectangle{}
		s.Area()
		s.Perimeter()

		var s2 Shape = &Circle{}
		s2.Area()
		s2.Perimeter()*/

	emp := &Employee{
		Person: Person{
			Name: "zhangsan",
			Age:  18,
		},
		EmployeeID: 12343,
	}
	emp.PrintInfo()
}

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct{}

func (*Rectangle) Area() {
	fmt.Println("Rectangle implement Area function")
}
func (*Rectangle) Perimeter() {
	fmt.Println("Rectangle implement Perimeter function")
}

type Circle struct{}

func (*Circle) Area() {
	fmt.Println("Circle implement Area function")
}
func (*Circle) Perimeter() {
	fmt.Println("Circle implement Perimeter function")
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name:%v,Age:%v,EmployeeID:%v \n", e.Name, e.Age, e.EmployeeID)
}
