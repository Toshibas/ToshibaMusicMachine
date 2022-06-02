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

	fmt.Println(dbConnector.guildExists("qwertyuiop"))
	fmt.Println(dbConnector.guildExists("qwertyuioq"))

	dbConnector.closeConnection()

}
