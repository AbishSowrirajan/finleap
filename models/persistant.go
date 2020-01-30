package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// CityJSON  data to be stored in DB
type CityJSON struct {
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

//DbLayer interface implementation for all the Db operations
type DbLayer interface {
	Create(CityJSON) interface{}
}

// ModelError for creating custom error ...
type ModelError struct {
	ErrCode string
	Err     error
	ErrTyp  string
}

// Create will create new city to the table
func (dbdata CityJSON) Create(data CityJSON) interface{} {

	var customerror ModelError

	err := Db.QueryRow("SELECT CITYNAME,CITYLONG,CITYLATI) FROM Finleap.FLCITY WHERE CITYNAME = ? GROUP BY CITYNAME,CITYLONG,CITYLATI ORDER BY CITYNAME", data.Name).Scan(&data.Name, &data.Longitude, &data.Latitude)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query City Table %w ", err.Error())
		customerror.ErrCode = "S1AUT901"
		customerror.ErrTyp = "500"

		return customerror
	}

	stmt, err := Db.Prepare("INSERT IGNORE INTO Finleap.FLCITY (CITYNAME,CITYLONG,CITYLATI)  VALUES(?,?,?)")
	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Customer Table %w ", err.Error())
		customerror.ErrCode = "S1AUT903"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(data.Name, data.Longitude, data.Latitude)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Customer Table %w ", err.Error())
		customerror.ErrCode = "S1AUT904"
		customerror.ErrTyp = "500"

		return customerror
	}

	return nil

}
