package main

import (
	"flag"
	"log"
	c "roboInsert/config"
	"roboInsert/sql"
)

var createConfig bool

func main() {
	flag.BoolVar(&createConfig, "c", false, "create config.yaml file")
	flag.Parse()

	if createConfig {
		c.CreateConfigFile()
		return
	}

	log.Print("loading config file")
	if err := c.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	log.Print("connecting sql ...")
	connection, err := sql.MakeSQL(c.Config.SQL.Host, c.Config.SQL.Port, c.Config.SQL.User, c.Config.SQL.Password)
	if err != nil {
		log.Println(err)
		return
	}
	sql.SetSQLConn(connection)
	connection.SearchClient()

}
