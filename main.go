package main

import (
	"demo/password/account"
	"demo/password/cloud"
	"demo/password/output"

	"fmt"

	"github.com/fatih/color"
)

func main() {
	var choice int
	// vault := account.NewVault(files.NewJsonDb("JSONbase.json"))
	vault := account.NewVault(cloud.NewCloudDB("htttp://a.ru"))
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
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		case 4:
			break Main
		}
	}
}

func findAccount(vault *account.VaultWithDB) {
	url := promptData("Введите url: ")
	accounts := vault.FindAcc(url)
	if len(accounts) == 0 {
		output.Output("Акаунтов не найдено")
	}
	for _, acc := range accounts {
		acc.Output()
	}

}
func deleteAccount(vault *account.VaultWithDB) {
	url := promptData("Введите url: ")
	isDeleted := vault.DeleteAcc(url)
	if isDeleted {
		color.Green("Аккаунт удален")
	} else {
		output.Output("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.Output("Неверный формат URL или Login")
		return
	}
	fmt.Println(myAccount)
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	color.Blue(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
