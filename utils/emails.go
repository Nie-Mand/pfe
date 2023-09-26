package utils

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)


type Destination struct {
    Email string  `csv:"email"`
    Company  string  `csv:"company"`
    Position string `csv:"position"`
}

func LoadEmails(emailsFile string) []Destination {

	f, err := os.Open(emailsFile)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _destinations := []*Destination{}

	if err := gocsv.UnmarshalFile(f, &_destinations); err != nil { 
		panic(err)
	}

    data := []Destination{}
    for _, d := range _destinations {
        data = append(data, *d)
	}

    return data
}