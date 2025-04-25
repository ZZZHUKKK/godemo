package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAT time.Time `json:"updatedAt"`
}

func (vault *Vault) ToByte() ([]byte, error) {
	data, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (vault *Vault) FindAcc(urlAcc string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, urlAcc) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAcc(urlAcc string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, urlAcc)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.UpdatedAT = time.Now()
	data, err := vault.ToByte()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "JSONbase.json")
	return isDeleted
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAT = time.Now()
	data, err := vault.ToByte()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "JSONbase.json")
}

func NewVault() *Vault {
	file, err := files.ReadFile("JSONbase.json")
	if err != nil {
		preVault := &Vault{
			Accounts:  []Account{},
			UpdatedAT: time.Now(),
		}
		return preVault
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}
