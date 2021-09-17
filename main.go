package main

import (
	"flag"
	"log"
)

var createConfig bool

const apiUrl = "https://maps.googleapis.com/maps/api/geocode/json"
const key = "AIzaSyDybcJ7PHZPP2es7YN_hd0D5OWjIR3kue0"

const raio float64 = 6367

func main() {
	flag.BoolVar(&createConfig, "c", false, "create config.yaml file")
	flag.Parse()

	if createConfig {
		CreateConfigFile()
		return
	}

	log.Print("loading config file")
	if err := LoadConfig(); err != nil {
		log.Fatal(err)
	}

	log.Print("connecting sql ...")
	connection, err := MakeSQL(Config.SQL.Host, Config.SQL.Port, Config.SQL.User, Config.SQL.Password)
	if err != nil {
		log.Println(err)
		return
	}
	SetSQLConn(connection)
	connection.SearchClient()

}
