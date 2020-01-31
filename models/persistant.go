package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// CityJSON  data to be stored in DB
type CityJSON struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// TempJSON  data to be stored in DB
type TempJSON struct {
	ID        string  `json:"id"`
	CityID    string  `json:"city_id"`
	Max       float64 `json:"max"`
	Min       float64 `json:"min"`
	Timestamp string  `json:"timestamp"`
}

// ForcastJSON  data to be stored in DB
type ForcastJSON struct {
	CityID string `json:"city_id"`
	Max    string `json:"max"`
	Min    string `json:"min"`
	Sample string `json:"sample"`
}

// WebHooksJSON ...
type WebHooksJSON struct {
	ID          string `json:"id"`
	CityID      string `json:"city_id"`
	CallbackURL string `json:"callback_url"`
}

// CompleteData ...
type CompleteData struct {
	CityJSON
	TempJSON
	ForcastJSON
}

//DbLayer interface implementation for all the Db operations
type DbLayer interface {
	Create(CityJSON) interface{}
	GetCity(CityJSON) interface{}
	UpdateCity(CityJSON) interface{}
	DeleteCity(CityJSON) interface{}
	InsertCityTemp(TempJSON) interface{}
	GetCityByName(CityJSON) interface{}
	GetTempByTimSt(TempJSON) interface{}
	GetAvgTempByCity(ForcastJSON) interface{}
	InsertCityWebH(WebHooksJSON) interface{}
	GetWebH(WebHooksJSON) interface{}
	DeleteWebH(WebHooksJSON) interface{}
}

// ModelError for creating custom error ...
type ModelError struct {
	ErrCode string `json :"Errcode"`
	Err     string `json:"Errormessage"`
	ErrTyp  string `json:"ErrorType"`
}

// Create will create new city to the table
func (dbdata CityJSON) Create(data CityJSON) interface{} {

	var customerror ModelError

	var cityExist bool

	var id int

	err := Db.QueryRow("SELECT CITYNAME,CITYLONG,CITYLATI ,IF(COUNT(*),'true','false') , LAST_INSERT_ID()  FROM Finleap.FLCITY WHERE CITYNAME = ? GROUP BY CITYNAME,CITYLONG,CITYLATI ORDER BY CITYNAME", data.Name).Scan(&data.Name, &data.Longitude, &data.Latitude, &cityExist, &id)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query City Table %w ", err).Error()
		customerror.ErrCode = "100"
		customerror.ErrTyp = "500"

		return customerror
	}

	if cityExist {

		customerror.Err = "City already exist"
		customerror.ErrCode = "101"
		customerror.ErrTyp = "400"

		return customerror

	}

	stmt, err := Db.Prepare("INSERT IGNORE INTO Finleap.FLCITY (CITYNAME,CITYLONG,CITYLATI)  VALUES(?,?,?)")
	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting City Table %w ", err).Error()
		customerror.ErrCode = "102"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(data.Name, data.Longitude, data.Latitude)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting City Table %w ", err).Error()
		customerror.ErrCode = "103"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// GetCity will get city details  from the  table
func (dbdata CityJSON) GetCity(data CityJSON) interface{} {

	var customerror ModelError

	var status int

	err := Db.QueryRow("SELECT CITYID , CITYNAME,CITYLONG,CITYLATI,CITYSTATUS  FROM Finleap.FLCITY WHERE CITYID = ? GROUP BY CITYID, CITYNAME,CITYLONG,CITYLATI,CITYSTATUS ORDER BY CITYNAME", data.ID).Scan(&data.ID, &data.Name, &data.Longitude, &data.Latitude, &status)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query City Table %w ", err).Error()
		customerror.ErrCode = "104"
		customerror.ErrTyp = "500"

		return customerror
	}

	if err == sql.ErrNoRows || status == 1 {

		customerror.Err = "City doesnt exist"
		customerror.ErrCode = "105"
		customerror.ErrTyp = "400"

		return customerror

	}

	return data

}

// GetCityByName will get city details  from the  table
func (dbdata CityJSON) GetCityByName(data CityJSON) interface{} {

	var customerror ModelError

	var status int

	err := Db.QueryRow("SELECT CITYID , CITYNAME,CITYLONG,CITYLATI,CITYSTATUS  FROM Finleap.FLCITY WHERE CITYNAME = ? GROUP BY CITYID, CITYNAME,CITYLONG,CITYLATI,CITYSTATUS ORDER BY CITYNAME", data.Name).Scan(&data.ID, &data.Name, &data.Longitude, &data.Latitude, &status)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query City Table %w ", err).Error()
		customerror.ErrCode = "106"
		customerror.ErrTyp = "500"

		return customerror
	}

	if err == sql.ErrNoRows || status == 1 {

		customerror.Err = "City doesnt exist"
		customerror.ErrCode = "107"
		customerror.ErrTyp = "400"

		return customerror

	}

	return data

}

// UpdateCity will update city to the table
func (dbdata CityJSON) UpdateCity(data CityJSON) interface{} {

	var customerror ModelError

	stmt, err := Db.Prepare("UPDATE Finleap.FLCITY SET CITYNAME = ? ,CITYLONG = ? ,CITYLATI = ? WHERE CITYID = ?")
	if err != nil {

		customerror.Err = fmt.Errorf("Error While Updating City Table %w ", err).Error()
		customerror.ErrCode = "108"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(data.Name, data.Longitude, data.Latitude, data.ID)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While updating city  Table %w ", err).Error()
		customerror.ErrCode = "109"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// DeleteCity will delete city from  the table
func (dbdata CityJSON) DeleteCity(data CityJSON) interface{} {

	var customerror ModelError

	stmt, err := Db.Prepare("UPDATE Finleap.FLCITY SET CITYSTATUS = ?  WHERE CITYID = ?")
	if err != nil {

		customerror.Err = fmt.Errorf("Error While Updating City Table %w ", err).Error()
		customerror.ErrCode = "110"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(1, data.ID)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While updating city  Table %w ", err).Error()
		customerror.ErrCode = "111"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// InsertCityTemp will create new city to the table
func (dbdata CityJSON) InsertCityTemp(data TempJSON) interface{} {

	var customerror ModelError

	stmt, err := Db.Prepare("INSERT IGNORE INTO Finleap.FLTEMP (TEMPCITYID,TEMPMAX,TEMPMIN,TEMPTIMEST)  VALUES(?,?,?,?)")

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Temperature Table %w ", err).Error()
		customerror.ErrCode = "112"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(data.CityID, data.Max, data.Min, data.Timestamp)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Temperature Table %w ", err).Error()
		customerror.ErrCode = "113"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// GetTempByTimSt will get city  temperature details  from the  table
func (dbdata CityJSON) GetTempByTimSt(data TempJSON) interface{} {

	var customerror ModelError

	err := Db.QueryRow("SELECT TEMPID , TEMPCITYID,TEMPMAX,TEMPMIN,TEMPTIMEST  FROM Finleap.FLTEMP WHERE TEMPTIMEST = ? AND TEMPCITYID = ? GROUP BY TEMPID,TEMPCITYID,TEMPMAX,TEMPMIN,TEMPTIMEST ORDER BY TEMPTIMEST", data.Timestamp, data.CityID).Scan(&data.ID, &data.CityID, &data.Max, &data.Min, &data.Timestamp)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query Temperature Table %w ", err).Error()
		customerror.ErrCode = "114"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// GetAvgTempByCity will get city  temperature details  from the  table
func (dbdata CityJSON) GetAvgTempByCity(data ForcastJSON) interface{} {

	var customerror ModelError

	rows, err := Db.Query("SELECT COUNT(TEMPCITYID), AVG(TEMPMAX) AS MAX , AVG(TEMPMIN) AS MIN  FROM Finleap.FLTEMP WHERE TEMPCITYID = ?", data.CityID)

	for rows.Next() {
		if err := rows.Scan(&data.Sample, &data.Max, &data.Min); err != nil {
			log.Fatal(err.Error())
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query Temperature Table %w ", err).Error()
		customerror.ErrCode = "115"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// InsertCityWebH will create new city to the table
func (dbdata CityJSON) InsertCityWebH(data WebHooksJSON) interface{} {

	var customerror ModelError

	var UrlExist bool

	err := Db.QueryRow("SELECT WEBHCITYID , IF(COUNT(*),'true','false') FROM Finleap.FLWEBH WHERE WEBHCITYID = ? AND WEBHURL = ? GROUP BY WEBHCITYID ORDER BY WEBHCITYID", data.CityID, data.CallbackURL).Scan(&data.CityID, &UrlExist)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query WebHooks Table %w ", err).Error()
		customerror.ErrCode = "116"
		customerror.ErrTyp = "500"

		return customerror
	}

	if UrlExist {

		customerror.Err = "Url  already exist for this city "
		customerror.ErrCode = "117"
		customerror.ErrTyp = "400"

		return customerror

	}

	stmt, err := Db.Prepare("INSERT IGNORE INTO Finleap.FLWEBH (WEBHCITYID,WEBHURL)  VALUES(?,?)")

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Webhooks Table %w ", err).Error()
		customerror.ErrCode = "118"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(data.CityID, data.CallbackURL)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While inseting Temperature Table %w ", err).Error()
		customerror.ErrCode = "119"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}

// GetWebH will get city  temperature details  from the  table
func (dbdata CityJSON) GetWebH(data WebHooksJSON) interface{} {

	var customerror ModelError

	err := Db.QueryRow("SELECT WEBHID , WEBHCITYID , WEBHURL  FROM Finleap.FLWEBH WHERE WEBHCITYID = ? AND WEBHURL = ? GROUP BY WEBHID , WEBHCITYID , WEBHURL ORDER BY WEBHCITYID", data.CityID, data.CallbackURL).Scan(&data.ID, &data.CityID, &data.CallbackURL)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query WebHooks Table %w ", err).Error()
		customerror.ErrCode = "120"
		customerror.ErrTyp = "500"

		return customerror
	}

	if err == sql.ErrNoRows {

		customerror.Err = "WebHooks Doesnt Exist"
		customerror.ErrCode = "121"
		customerror.ErrTyp = "400"

		return customerror
	}

	return data

}

// DeleteWebH will delete city from  the table
func (dbdata CityJSON) DeleteWebH(data WebHooksJSON) interface{} {

	var customerror ModelError

	err := Db.QueryRow("SELECT WEBHID , WEBHCITYID , WEBHURL  FROM Finleap.FLWEBH WHERE WEBHID = ? GROUP BY WEBHID , WEBHCITYID , WEBHURL ORDER BY WEBHID", data.ID).Scan(&data.ID, &data.CityID, &data.CallbackURL)

	if err != nil && err != sql.ErrNoRows {

		customerror.Err = fmt.Errorf("Error While Query WebHooks Table %w ", err).Error()
		customerror.ErrCode = "122"
		customerror.ErrTyp = "500"

		return customerror
	}

	if err == sql.ErrNoRows {

		customerror.Err = "WebHooks Doesnt Exist"
		customerror.ErrCode = "123"
		customerror.ErrTyp = "400"

		return customerror
	}

	stmt, err := Db.Prepare("UPDATE Finleap.FLWEBH SET WEBHSTATUS = ?  WHERE WEBHID = ?")
	if err != nil {

		customerror.Err = fmt.Errorf("Error While Deleting Webhook  Table %w ", err).Error()
		customerror.ErrCode = "124"
		customerror.ErrTyp = "500"

		return customerror
	}
	defer stmt.Close()

	stmt.Exec(1, data.ID)

	if err != nil {

		customerror.Err = fmt.Errorf("Error While Deleting Web Hook  city  Table %w ", err).Error()
		customerror.ErrCode = "125"
		customerror.ErrTyp = "500"

		return customerror
	}

	return data

}
