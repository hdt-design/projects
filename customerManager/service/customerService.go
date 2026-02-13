// Package service
package service

import (
	"main/shangguigu/customerManager/model"
)

// CustomerService 完成对 Customer 的操作，
// 包括增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回*CustomerService

func NewCustomerService() *CustomerService {
	//为了能够看到客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片

func (c *CustomerService) List() []model.Customer {
	return c.customers
}

//新增客户
//!!!

func (c *CustomerService) Add(cusomer model.Customer) bool {

	//确定一个分配id的规则，就是添加的顺序
	c.customerNum++
	cusomer.ID = c.customerNum
	c.customers = append(c.customers, cusomer)
	return true
}

//根据id删除用户

func (c *CustomerService) Delete(id int) bool {
	index := c.FindByID(id)
	if index == -1 {
		return false
	}
	//删除切片中的元素
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

//根据id查找客户在切片中对应下标，如果无该用户返回-1

func (c *CustomerService) FindByID(id int) int {
	index := -1
	//遍历c.customers切片，找到id对应的下标
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].ID == id {
			index = i
		}
	}
	return index
}
