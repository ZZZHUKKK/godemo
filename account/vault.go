package account

import (
	"demo/password/encrypt"
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type VaultWithDB struct {
	Vault
	db  Db
	enc encrypt.Encrypt
}

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

func (vault *VaultWithDB) FindAcc(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		if checker(account, str) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDB) DeleteAcc(urlAcc string) bool {
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
	data, err := vault.Vault.ToByte()
	if err != nil {
		output.Output(err)
	}
	vault.db.Write(data)
	return isDeleted
}

func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAT = time.Now()
	data, err := vault.Vault.ToByte()
	if err != nil {
		output.Output(err.Error())
	}
	vault.db.Write(data)
}

func NewVault(db Db, enc encrypt.Encrypt) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		preVault := &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAT: time.Now(),
			},
			db:  db,
			enc: enc,
		}
		return preVault
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red(err.Error())
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAT: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}
