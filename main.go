package main

import (
	"demo/password/account"

	"fmt"

	"github.com/fatih/color"
)

func main() {

	login := promptData("Введите логин")

	password := promptData("Введите пароль")

	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTime(login, password, url)

	if err != nil {

		fmt.Println("Неверный формат URL или Login")

		return

	}

	myAccount.Output()

	fmt.Println(myAccount)

}

func promptData(prompt string) string {

	color.Blue(prompt + ": ")

	var res string

	fmt.Scan(&res)

	return res

}
