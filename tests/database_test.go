package tests

import (
	database "../my_database"
	"testing"
	"bufio"
	"strings"
)





func TestInsert(t *testing.T){
	dataStorage := database.MyDatabaseWrapper{}
	dataStorage.Initialize()
	fields := []string{"a", "b", "c"}
	dataStorage.AddOneRecord(fields)

}