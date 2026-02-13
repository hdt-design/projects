// Package model
package model

import "fmt"

//声明一个Customer结构体，表示客户信息

type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个新的Customer实例

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建Customer实例方法，不带id

func NewCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户信息，格式化字符串

func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		c.ID, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}
