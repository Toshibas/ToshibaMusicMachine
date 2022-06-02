package main

import (

	"log"
	"fmt"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"

)

type DBConnector struct {
	address string
	port string
	user string
	password string
	database string
	db *sql.DB 
}

func createDBConnector(address, port, user, password, database string) (DBConnector, error) {

	connectionRow := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, database)

	db, err := sql.Open("mysql", connectionRow)

	if err != nil {
		return DBConnector{}, err
	}

	connector := DBConnector{
		address: address,
		port: port,
		user: user,
		password: password,
		database: database,
		db: db,
	}
	
	return connector, nil
}

func (connector DBConnector) closeConnection() {

	connector.db.Close()

}

func (connector DBConnector) guildExists(guildID string) bool {

	rows, err := connector.db.Query("SELECT * FROM guilds WHERE id = ?", guildID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	return rows.Next() 

}
