package files

import (
	"demo/password/output"
	"fmt"
	"os"
)

type JsonDb struct {
	Filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		Filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.Filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.Filename)
	if err != nil {
		output.Output(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.Output(err)
		return
	}
	fmt.Println("Запись успешна")
}
