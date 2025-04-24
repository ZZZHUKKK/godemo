package main

import (
	"demo/password/account"
	"demo/password/files"

	"fmt"

	"github.com/fatih/color"
)

func main() {
	var choice int
Main:
	for {
		fmt.Println(`Введите цифру:
1: Создать аккаунт
2: Найти аккаунт
3: Удалить аккаунт
4: Exit`)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			createAccount()
		case 2:
			foundAccount()
		case 3:
			deleteAccount()
		case 4:
			break Main
		}
	}
}

func foundAccount()  {}
func deleteAccount() {}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}
	fmt.Println(myAccount)
	byteAccount, err := myAccount.ToByte()
	if err != nil {
		fmt.Println("Неверный формат JSON")
		return
	}
	files.WriteFile(byteAccount, "JSONbase")
}

func promptData(prompt string) string {
	color.Blue(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
