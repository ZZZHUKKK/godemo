package account

import (
	"errors"

	// "reflect"

	"fmt"

	"math/rand/v2"

	"net/url"

	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "1" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	preAccount := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if password == "1" {
		preAccount.generatePassword(12)
	}
	// field, _ := reflect.TypeOf(preAccount).Elem().FieldByName("login")
	// fmt.Println(string(field.Tag))
	return preAccount, nil
}

// func newAccount(login, password, urlString string) (*account, error) {

// if login == "1" {

// return nil, errors.New("INVALID_LOGIN")

// }

// _, err := url.ParseRequestURI(urlString)

// if err != nil {

// return nil, errors.New("INVALID_URL")

// }

// preAccount := &account{

// url: urlString,

// login: login,

// password: password,

// }

// if password == "1" {

// preAccount.generatePassword(12)

// }

// return preAccount, nil

// }
