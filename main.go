package main

import (
	"demo/password/account"
	"demo/password/encrypt"
	"demo/password/files"
	"demo/password/output"
	"strings"

	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menuMap = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountUrl,
	"3": findAccountLogin,
	"4": deleteAccount,
}

// func menuCounter() func() {
// 	i := 0
// 	return func() {
// 		i++
// 		fmt.Println(i)
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading .env file")
	}
	var choice string
	enc, _ := encrypt.NewEncrypt()
	vault := account.NewVault(files.NewJsonDb("data.vault"), *enc)
	//vault := account.NewVault(cloud.NewCloudDB("htttp://a.ru"))
	// counter := menuCounter()
Main:
	for {
		// counter()
		menu := []string{"1: Создать аккаунт", "2: Найти аккаунт по URL", "3: Найти аккаунт по Login", "4: Удалить аккаунт", "5: Exit", "Введите цифру"}
		choice = promptData(menu)
		menuFunc := menuMap[choice]
		if menuFunc == nil {
			break Main
		}
		menuFunc(vault)
		// switch choice {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// case "4":
		// 	break Main
		// }
	}
}

func findAccountUrl(vault *account.VaultWithDB) {
	url := promptData("Введите url: ")
	accounts := vault.FindAcc(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputRes(&accounts)
}

func findAccountLogin(vault *account.VaultWithDB) {
	login := promptData("Введите login: ")
	accounts := vault.FindAcc(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputRes(&accounts)
}

func outputRes(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.Output("Акаунтов не найдено")
	}
	for _, acc := range *accounts {
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

func promptData(prompt ...any) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}
