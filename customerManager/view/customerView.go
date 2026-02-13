package main

import (
	"fmt"
	"main/shangguigu/customerManager/model"
	"main/shangguigu/customerManager/service"
)

type customerView struct {

	//定义必要字段
	key  string //接收用户输入
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有客户信息

func (c *customerView) list() {

	//获取当前所有客户信息
	customers := c.customerService.List()
	//显示
	fmt.Println("------------------------客户列表-----------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].ID,"\t",customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------客户列表完成-------------------------")
}

//得到用户的输入，信息构建新的用户，并完成添加

func (c *customerView) add() {
	fmt.Println("------------------------添加用户------------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意：id不让用户输入，是系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if c.customerService.Add(customer) {
		fmt.Println("添加成功！")
	} else {
		fmt.Println("添加失败！")
	}
}

//得到用户的输入id，删除id对应的用户

func (c *customerView) delect() {
	fmt.Println("------------------------删除用户--------------------------------")
	fmt.Println("请选择待删除客户编号（-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)")
	//这里可以加入循环判断
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		//调用删除方法
		if c.customerService.Delete(id) {
			fmt.Println("删除成功！")
		} else {
			fmt.Println("删除失败！")
		}
	}
}

//退出

func (c *customerView) exit() {

	fmt.Println("确认是否退出(Y/N)")
	for {
		fmt.Scanln(&c.key)
		if c.key == "Y" || c.key == "y" {
			c.loop = false
			break
		} else if c.key == "N" || c.key == "n" {
			c.loop = true
			break
		} else {
			fmt.Println("输入错误，请重新输入！")
		}
	}
}

//显示主菜单

func (c *customerView) mainMenu() {
	for {

		fmt.Println("----------------客户信息管理--------------")
		fmt.Println("                1.添加用户")
		fmt.Println("                2.修改用户")
		fmt.Println("                3.删除用户")
		fmt.Println("                4.客户列表")
		fmt.Println("                5.退出")
		fmt.Println("请选择1-5")

		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			c.add()
		case "2":
			fmt.Println("功能暂未开通！")
		case "3":
			c.delect()
		case "4":
			c.list()
		case "5":
			c.exit()
		default:
			fmt.Println("输入错误，请重新输入！")
		}
		if !c.loop {
			break
		}
	}
	fmt.Println("欢迎下次光临！")
}

func main() {
	//创建一个customerView
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//创建一个customerService
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
