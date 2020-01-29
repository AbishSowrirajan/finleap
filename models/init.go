package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	con := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(mysql:" + strconv.Itoa(cnfg.Port) + ")/Finleap_Weather?charset=utf8"
	fmt.Println(con)
	_, err := sql.Open("mysql", con)

	if err != nil {

		log.Fatalf("Error in connecting to Database  %v ", err.Error())
	}

	log.Println("DB connected")
}
