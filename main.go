package main

import (
	
	"fmt"
	"log"

)

 
func main() {

	dbConnector, err := createDBConnector("54.226.233.122", "3306", "bodko", "1234", "toshibaMusicMachine")

	if err != nil {
		log.Fatal(err)
	}

	// dbConnector.addGuild("111111", true)

	fmt.Println(dbConnector.getGuildByID("111112"))

	// fmt.Println(dbConnector.guildExists("111111"))
	

	dbConnector.closeConnection()

}
