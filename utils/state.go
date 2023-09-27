package utils

import (
	"log"
	"os"
	"strings"
)


const (
	STATE_FILE = ".pfe.state"
)

var hashes = make(map[string]bool)
var newHashed = []string{}


func ReadStateFile() map[string]bool {
	f, err := os.ReadFile(STATE_FILE)
	
	if err != nil {
		log.Fatal(err)
	}

	text := string(f)

	hashes := map[string]bool{}

	if text != "" {
		for _, hash := range strings.Split(text, "\n") {
			hashes[hash] = true
		}
	}

	return hashes
}

func CreateStateFile() {
	f, err := os.Create(STATE_FILE) 

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

func init() {
	if _, err := os.Stat(STATE_FILE); os.IsNotExist(err) {
		CreateStateFile()
	} else {
		hashes = ReadStateFile()
	}
}

func FilterRows(emails []Destination) []Destination {
	var filtered []Destination

	for _, email := range emails {
		if _, ok := hashes[email.Hash()]; !ok {
			filtered = append(filtered, email)
		}
	}

	return filtered
} 

func AddHash(destination Destination) {
	newHashed = append(newHashed, destination.Hash())
}

func SaveNewHashes() {
	f, err := os.OpenFile(STATE_FILE, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, hash := range newHashed {
		if _, err := f.WriteString(hash + "\n"); err != nil {
			log.Fatal(err)
		}
	}
}