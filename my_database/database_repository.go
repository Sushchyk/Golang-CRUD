package my_database

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"bufio"
	"fmt"
)

type MyDatabaseRepository struct {
	filename string
}

func (databaseRepo *MyDatabaseRepository)load() map[int]Employee {

	var result map[int]Employee
	data, err := ioutil.ReadFile(databaseRepo.filename)
	if (err == nil) {
		json.Unmarshal(data, &result)
	} else {
		return make(map[int]Employee)
	}
	return result
}

func (databaseRepo *MyDatabaseRepository)save(data map[int]Employee) error {
	fileHandle, _ := os.Create(databaseRepo.filename)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	b, err := json.Marshal(data)

	if (err == nil) {
		fmt.Fprintln(writer, string(b))
	}

	writer.Flush()
	return nil
}