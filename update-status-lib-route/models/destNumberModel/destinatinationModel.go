package destnumbermodel

import (
	"update-status/config"
	"update-status/entities"
)

func UpdateStatus(dnumber string, destnumber entities.DestNumber) bool {
	query, err := config.DB.Exec(`UPDATE getroutev3 SET status = ? WHERE destination_number = ?`, destnumber.Status, dnumber)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}