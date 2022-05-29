package main

import (
	"log"
	"os/exec"
)

func main() {

	_, err := exec.Command("./scripts/convertToDCA.sh", "./qwe.mp3", "./qwe.dca").Output()
	if err != nil {
		log.Fatal(err)
	}

}
