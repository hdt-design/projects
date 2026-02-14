package service

import (
	"customer/model"
	"fmt"
)

type CustomerService struct {
	Customers []model.Customer
}

func NewCustomerService() *CustomerService {
	var service CustomerService
	service.Customers = append(service.Customers,
		model.NewCustomer("zs", "男", 18,
			"151", "zs151@qq.com"))
	return &service
}

func (c *CustomerService) List() []model.Customer {
	return c.Customers
}

func (c *CustomerService) Add(name string, gender string, age int,
	phone, email string) {
	c.Customers = append(c.Customers, model.NewCustomer(name, gender, age, phone, email))
}

func (c *CustomerService) FindId(id int) int {
	index := -1
	for i := 0; i < len(c.Customers); i++ {
		if c.Customers[i].Id == id {
			index = i
			break
		}
	}
	return index
}

func (c *CustomerService) Delect(id int) bool {
	index := c.FindId(id)
	if index == -1 {
		fmt.Println("该用户不存在")
		return false
	}
	c.Customers = append(c.Customers[:index], c.Customers[index+1:]...)
	return true
}
