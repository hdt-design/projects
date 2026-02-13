// Package utils
package utils

import "fmt"

//定义一个结构体FamilyAccount，用于记录家庭账户信息

type FamilyAccount struct {

	//声明必须的字段

	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个字段，控制是否退出for
	loop bool
	//定义账户余额
	balance float64
	//收支金额
	money float64
	//说明
	note string
	//定义一个字段，记录是否有收支
	flag bool
	//收支详情使用string记录
	//有收支时对details拼接
	details string
}

//编写要给工厂模式的构造方法，返回一个*FamilyAccount实例

func NewFamilyAccount() *FamilyAccount {

	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说    明",
	}
}

//将显示明细写成一个方法

func (f *FamilyAccount) ShowDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if f.flag {
		fmt.Println(f.details)
	} else {
		fmt.Println("暂无收支记录")
	}
}

//将登记收入写成一个方法，和*FamilyAccount绑定

func (f *FamilyAccount) income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("本次收入说明")
	fmt.Scanln(&f.note)
	f.flag = true
	f.details += fmt.Sprintf("\n收入\t%.2f\t%.2f\t%s", f.balance, f.money, f.note)
}

//将登记支出写成一个方法，和*FamilyAccount绑定

func (f *FamilyAccount) pay() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&f.money)
	//判断
	if f.money > f.balance {
		fmt.Println("余额不足，请重新输入")
	}
	f.balance -= f.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&f.note)
	f.flag = true
	f.details += fmt.Sprintf("\n支出\t%.2f\t%.2f\t%s",
		f.balance, f.money, f.note)
}

//将退出系统写成一个方法,和*FamilyAccount绑定

func (f *FamilyAccount) exit() {

	fmt.Println("确定退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N" {
			break
		} else {
			fmt.Println("输入错误，请重新输入")
		}
	}
	if choice == "y" || choice == "Y" {
		f.loop = false
	}
}

//给该结构体绑定相应的方法
//显示主菜单

func (f *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n-----------------家庭收支记账软件----------------")
		fmt.Println("                    1. 收入")
		fmt.Println("                    2. 支出")
		fmt.Println("                    3. 显示明细")
		fmt.Println("                    4. 退出系统")
		fmt.Scanln(&f.key)
		switch f.key {
		case "1":
			f.income()
		case "2":
			f.pay()
		case "3":
			f.ShowDetails()
		case "4":
			f.exit()
		default:
			fmt.Println("输入错误，请重新输入")
		}
		if !f.loop {
			break
		}
	}
}
