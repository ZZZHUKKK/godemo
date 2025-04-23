package account

import (
	"errors"

	"fmt"

	"math/rand/v2"

	"net/url"

	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login string

	password string

	url string
}

type accountWithTimestamp struct {
	createdAt time.Time

	updatedAt time.Time

	Account
}

func (acc *accountWithTimestamp) Output() {

	fmt.Println(acc.login, acc.password, acc.url)

}

func (acc *accountWithTimestamp) generatePassword(n int) {

	res := make([]rune, n)

	for i := range res {

		res[i] = letterRunes[rand.IntN(len(letterRunes))]

	}

	acc.password = string(res)

}

func NewAccountWithTime(login, password, urlString string) (*accountWithTimestamp, error) {

	if login == "1" {

		return nil, errors.New("INVALID_LOGIN")

	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {

		return nil, errors.New("INVALID_URL")

	}

	preAccount := &accountWithTimestamp{

		createdAt: time.Now(),

		updatedAt: time.Now(),

		Account: Account{

			url: urlString,

			login: login,

			password: password,
		},
	}

	if password == "1" {

		preAccount.generatePassword(12)

	}

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
