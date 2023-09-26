package utils

import (
	"fmt"
	"log"
	"os"
)


func GetArtwork() {
	file, err := os.ReadFile("ART")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(file))
}