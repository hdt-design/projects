package main

import (
	"bufio"
	"customer/service"
	"customer/toolfunc"
	"fmt"
	"os"
	"strings"
)

type customerView struct {
	loop          bool
	key           int
	customService *service.CustomerService
}

func newCustomView() *customerView {
	return &customerView{
		loop:          true,
		key:           0,
		customService: service.NewCustomerService(),
	}
}

func (c *customerView) add() {
	var name string
	fmt.Println("name")
	if _, err := fmt.Scanln(&name); err != nil {
		fmt.Println(err)
	}
	var gender string
	fmt.Println("gender")
	if _, err := fmt.Scanln(&gender); err != nil {
		fmt.Println(err)
	}
	var age int
	fmt.Println("age")
	if _, err := fmt.Scanln(&age); err != nil {
		fmt.Println(err)
	}
	var phone string
	fmt.Println("phone")
	if _, err := fmt.Scanln(&phone); err != nil {
		fmt.Println(err)
	}
	var email string
	fmt.Println("email")
	if _, err := fmt.Scanln(&email); err != nil {
		fmt.Println(err)
	}
	c.customService.Add(name, gender, age, phone, email)
	fmt.Println("---------------------添加成功---------------------")
}

func (c *customerView) delect() {
	var id int
	fmt.Println("请输入需要删除的id")
	if _, err := fmt.Scanln(&id); err != nil {
		fmt.Println(err)
	}
	var choice string
	fmt.Println("你确定要删除吗？Y/N")
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println(err)
	}
	for {
		if choice == "Y" || choice == "y" {
			c.customService.Delect(id)
			fmt.Println("----------------删除成功--------------------")
			break
		} else if choice == "N" || choice == "n" {
			break
		}
	}
}

func (c *customerView) reset() {
	var id int
	fmt.Println("请输入需要修改的id")
	if _, err := fmt.Scanln(&id); err != nil {
		fmt.Println(err)
	}
	index := c.customService.FindId(id)
	if index == -1 {
		fmt.Println("用户不存在")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		_, name, gender, age, phone, email := c.customService.List()[index].GetInfo()
		name = toolfunc.PromptString(scanner,
			fmt.Sprintf("姓名(%v)", name), name)
		gender = toolfunc.PromptString(scanner,
			fmt.Sprintf("性别(%v)", gender), gender)
		age = toolfunc.PromptIntWithDefault(scanner,
			fmt.Sprintf("年龄(%v)", age), age)
		phone = toolfunc.PromptString(scanner,
			fmt.Sprintf("电话(%v)", phone), phone)
		email = toolfunc.PromptString(scanner,
			fmt.Sprintf("邮箱(%v)", email), email)
		c.customService.List()[index].Reset(name, gender,
			age, phone, email)
	}
}

func (c *customerView) query() {
	fmt.Println("请选择查询方式（-1退出查询）：1，Id查询，2，姓名查询")
	var choice string
	if _, err := fmt.Scanln(&choice); err != nil {
		fmt.Println(err)
	}
	if choice == "-1" {
		return
	} else if choice == "1" {
		var id int
		fmt.Println("请输入需要查询的id")
		if _, err := fmt.Scanln(&id); err != nil {
			fmt.Println(err)
		}
		index := c.customService.FindId(id)
		if index == -1 {
			fmt.Println("该用户不存在")
		} else {
			fmt.Println("-----------客户信息----------------")
			fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
			a, b, c, d, e, f := c.customService.List()[index].GetInfo()
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t\t%v\n", a, b, c, d, e, f)
			fmt.Println("------------客户信息完成-------------")
			fmt.Println()
		}
	} else if choice == "2" {
		fmt.Println("选择查询方式：1.精准查询，2，模糊查询")
		var choice2 string
		if _, err := fmt.Scanln(&choice2); err != nil {
			fmt.Println(err)
		}
		var name string
		fmt.Println("请输入客户姓名")
		if _, err := fmt.Scanln(&name); err != nil {
			fmt.Println(err)
		}
		var customer = c.customService.List()
		if choice2 == "1" {
			var Isin = false
			for _, i := range customer {
				a, b, c, d, e, f := i.GetInfo()
				if b == name {
					Isin = true
					fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n",
						a, b, c, d, e, f)
				}
			}
			if Isin == false {
				fmt.Println("该用户不存在")
			}
		} else if choice2 == "2" {
			var Isin = false
			for _, i := range customer {
				a, b, c, d, e, f := i.GetInfo()
				if strings.Contains(b, name) {
					Isin = true
					fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n",
						a, b, c, d, e, f)
				}
			}
			if Isin == false {
				fmt.Println("该用户不存在")
			}
		}
	}
}

func (c *customerView) list() {
	var customer = c.customService.List()
	fmt.Println("--------------------客户列表----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, v := range customer {
		a, b, c, d, e, f := v.GetInfo()
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t\t%v\n", a, b, c, d, e, f)
		fmt.Println("---------------客户信息完成--------------")
	}
}

func (c *customerView) menu() {
	for {
		fmt.Println("----------------客户信息管理-------------")
		fmt.Println("-----------------1.添加用户--------------")
		fmt.Println("-----------------2.删除用户--------------")
		fmt.Println("-----------------3.修改用户--------------")
		fmt.Println("-----------------4.查询用户--------------")
		fmt.Println("-----------------5.用户列表--------------")
		fmt.Println("-----------------6.退出程序--------------")
		fmt.Println("请选择1-6")
		if _, err := fmt.Scanln(c.key); err != nil {
			fmt.Println(err)
		}
		switch c.key {
		case 1:
			c.add()
		case 2:
			c.delect()
		case 3:
			c.reset()
		case 4:
			c.query()
		case 5:
			c.list()
		case 6:
			for {
				var char string
				fmt.Println("确认退出吗？Y/N")
				if _, err := fmt.Scanln(&char); err != nil {
					fmt.Println(err)
				}
				if char == "Y" || char == "y" {
					c.loop = false
					break
				} else if char == "N" || char == "n" {
					break
				}
			}
		default:
			fmt.Println("请输入正确的数字")
		}
		if !c.loop {
			break
		}
	}
}

func main() {
	var c = newCustomView()
	c.menu()
}
