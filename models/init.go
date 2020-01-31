package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/AbishSowrirajan/finleap/config"
	_ "github.com/go-sql-driver/mysql"
)

// Db variable for Database operation
var Db *sql.DB

// Init for Database access
func Init() {

	var cnfg config.Cnfg

	cf, err := config.Config()

	if err != nil {

		log.Fatalf("Error in accessing configration file %v", err.Error())
	}

	_ = json.Unmarshal(cf, &cnfg)

	con := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(mysql:" + strconv.Itoa(cnfg.Port) + ")/Finleap?charset=utf8"
	fmt.Println(con)

	Db, err = sql.Open("mysql", con)

	if err != nil {

		log.Fatalf("Error in connecting to Database  %v ", err.Error())
	}

	log.Println("DB connected")
}
