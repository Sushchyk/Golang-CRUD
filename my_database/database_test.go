package my_database

import (
	"testing"
	"time"
	"strings"
)

//go test -coverprofile fmt

func TestInsert(t *testing.T){

	dataStorage := MyDatabaseWrapper{}
	dataStorage.Initialize()

	count_of_records := len(dataStorage.data)
	fields := []string{"a", "b", "c"}
	dataStorage.addOneRecord(fields)
	time.Sleep(500 * time.Millisecond)
	if count_of_records != len(dataStorage.data)-1{
		t.Error(
			"For", "Insert in db",
			"Expected", "Inserted",
			"got", "Error",
		)
	}

}


func TestUpdate(t *testing.T){

	dataStorage := MyDatabaseWrapper{}
	dataStorage.Initialize()

	fields := []string{"a1", "b1", "c1"}
	dataStorage.updateOneRecord(1, fields)
	time.Sleep(500 * time.Millisecond)
	updated := dataStorage.data[1]
	if updated.Name != fields[0] || updated.Surname != fields[1] || updated.Position != fields[2] {
		t.Error(
			"For", "Update in db",
			"Expected", "Updated",
			"got", "Error",
		)
	}

}

func TestDelete(t *testing.T){

	dataStorage := MyDatabaseWrapper{}
	dataStorage.Initialize()

	count_of_records := len(dataStorage.data)
	dataStorage.deleteOneRecord(1)
	time.Sleep(500 * time.Millisecond)
	if count_of_records != len(dataStorage.data)+1 {
		t.Error(
			"For", "Delete from db",
			"Expected", "Deleted",
			"got", "Error",
		)
	}
}

func TestValidateQuery(t *testing.T) {
	query_string := "CREATE v d a"
	query := strings.Fields(query_string)
	if ValidateQuery(query) == false {
		t.Error(
			"For", "Validate query",
			"Expected", "true",
			"got", "false",
		)
	}
}


func TestHandleQuery(t *testing.T) {
	dataStorage := MyDatabaseWrapper{}
	dataStorage.Initialize()

	query_string := "error a d d a da a"
	query := strings.Fields(query_string)

	if HandleQuery(query, &dataStorage) == "Wrong action" {
		t.Error(
			"For", "Handle query",
			"Expected", "OK",
			"got", "Wron action",
		)
	}
}