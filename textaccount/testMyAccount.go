package main

import (
	"fmt"
)

func main() {
	var key int
	var money float64
	var note string
	balance := 10000.00
	flag := false
	details := "收支\t账户金额\t收支金额\t说    明"

	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择功能：")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("-----------------当前收支明细记录----------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支，来一笔吧！")
			}
		case 2:
			fmt.Println("本次收入金额")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case 3:
			fmt.Println("本次支出金额")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足，请充值！")
			} else {
				balance -= money
				fmt.Println("本次支出说明")
				fmt.Scanln(&note)
				details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
				flag = true
			}
		case 4:
			fmt.Println("欢迎下次使用！")
			return
		}
	}
}
