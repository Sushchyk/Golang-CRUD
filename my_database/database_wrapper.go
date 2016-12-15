package my_database

import (
	"sync"
	"fmt"
	"errors"
)

type MyDatabaseWrapper struct {
	file sync.Mutex
	data map[int]Employee
	repo MyDatabaseRepository
}

func (databaseWrapper *MyDatabaseWrapper) loadDataFromFile() {
	databaseWrapper.data = databaseWrapper.repo.load()
}

func (databaseWrapper *MyDatabaseWrapper) saveDataToFile() {
	databaseWrapper.repo.save(databaseWrapper.data)
}

func (databaseWrapper *MyDatabaseWrapper) Initialize() {

	databaseWrapper.file = sync.Mutex{}
	databaseWrapper.repo = MyDatabaseRepository{filename:"test.json"}
	databaseWrapper.loadDataFromFile()

}

func (databaseWrapper *MyDatabaseWrapper) addOneRecord(fields []string) (error) {

	max_key := 0
	for k, _ := range databaseWrapper.data {
		if (k > max_key) {
			max_key = k
		}
	}
	databaseWrapper.data[max_key + 1] = Employee{Name: fields[0], Surname: fields[1], Position: fields[2]}

	databaseWrapper.saveDataToFile()

	return nil
}

func (databaseWrapper *MyDatabaseWrapper) updateOneRecord(id int, fields []string) (error) {

	if _, ok := databaseWrapper.data[id]; ok {
		databaseWrapper.data[id] = Employee{Name: fields[0], Surname: fields[1], Position: fields[2]}
		databaseWrapper.saveDataToFile()
		return nil
	} else {
		return errors.New("Record not found")
	}
}

func (databaseWrapper *MyDatabaseWrapper) readOneRecord(id int) (string, error) {
	if val, ok := databaseWrapper.data[id]; ok {
		return fmt.Sprintf("ID: %-3d || Name: %-12s || Surname: %-15s || Position: %-15s",
			id, val.Name, val.Surname, val.Position), nil
	} else {
		return "_", errors.New("Record not found")
	}
}

func (databaseWrapper *MyDatabaseWrapper) readAllRecords() (string, error) {
	result := fmt.Sprintf("%3s %15s %15s %15s", "ID", "Name", "Surname", "Position")

	for k, v := range databaseWrapper.data {
		result = result + fmt.Sprintf("\n%3d|%15s|%15s|%15s|",
			k, v.Name, v.Surname, v.Position)
	}

	return string(result), nil
}

func (databaseWrapper *MyDatabaseWrapper) deleteOneRecord(id int) error {
	if _, ok := databaseWrapper.data[id]; ok {
		delete(databaseWrapper.data, id)
		databaseWrapper.saveDataToFile()
		return nil
	} else {
		return errors.New("Record not found")
	}
}


