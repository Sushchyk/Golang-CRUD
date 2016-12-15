package my_database

import (
	"strconv"
	"errors"
)

//TODO: implement in other way!
func getActions() map[string]string {

	actions := map[string]string{
		"c" : "CREATE",
		"r" : "READ",
		"u" : "UPDATE",
		"d" : "DELETE",
	}

	return actions;
}



func ValidateQuery(query []string) bool {


	countOfFields := new (Employee).getCountOfFields()
	actions := getActions();

	if (len(query) == 0) {
		return false;
	}
	action := query[0]
	if (action != actions["c"] && action != actions["r"] && action != actions["u"] && action != actions["d"]) {
		return false;
	}

	if (action == actions["c"] && len(query) != (1 + countOfFields)) {
		return false;
	}

	if (action == actions["r"]) {
		if (len(query) != 1 && len(query) != 2) {
			return false
		}

		if (len(query) == 2 && !isNumber(query[1])) {
			return false
		}
	}

	if (action == actions["u"]) {
		if ((len(query) != 2 + countOfFields) || !isNumber(query[1])) {
			return false
		}
	}

	if (action == actions["d"]) {
		if ((len(query) != 2) || !isNumber(query[1])) {
			return false
		}
	}

	return true
}

func HandleQuery(query []string, dataStorage *MyDatabaseWrapper) string {
	actions := getActions()

	var responseMessage string
	var responseData string
	var err error

	switch query[0] {

	case actions["c"]:
		err = dataStorage.addOneRecord(query[1:4]);
		responseMessage = "Created"

	case actions["r"]:
		if (len(query) == 1) {
			responseData, err = dataStorage.readAllRecords();

		} else {
			id, typeErr := strconv.Atoi(query[1])
			if (typeErr == nil) {
				responseData, err = dataStorage.readOneRecord(id)
			}
		}


	case actions["u"]:
		id, typeErr := strconv.Atoi(query[1])
		if (typeErr == nil) {
			err = dataStorage.updateOneRecord(id, query[2:5])

		}	else {
			err = errors.New("ID must be integer")
		}

		responseMessage = "Updated"


	case actions["d"]:
		id, typeErr := strconv.Atoi(query[1])
		if (typeErr == nil) {
			err = dataStorage.deleteOneRecord(id)
		}

		responseMessage = "Deleted"
	}

	if (err == nil) {
		if (query[0] == actions["r"]) {
			return responseData
		}

		return responseMessage

	}	else {

		return err.Error()

	}
}