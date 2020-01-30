package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Cnfg parameters ......
type Cnfg struct {
	Port    int    `json:"Port"`
	Dbuname string `json:"Dbuname"`
	Dbpass  []byte `json:"Dbpass"`
}

//Port used for connecting to browser ...
type Port struct {
	DbPort   int `json:"dbport"`
	HostPort int `json:"hostport"`
}

// Config files are used to get the configraiton parameters ...
func Config() ([]byte, error) {

	file, err := os.Open("./config/config")

	dat, err := ioutil.ReadAll(file)

	if err != nil {

		return nil, fmt.Errorf("Error in reading Configration file : %w", err)
	}

	defer file.Close()

	var p Port

	var cf Cnfg

	err = json.Unmarshal(dat, &p)

	cf.Port = p.DbPort

	cf.Dbuname = os.Getenv("SQL_UNAME")

	cf.Dbpass = []byte(os.Getenv("SQL_PASS"))

	data, err := json.Marshal(cf)

	if err != nil {

		return nil, fmt.Errorf("Error in reading Configration file : %w", err)

	}

	return data, nil

}
