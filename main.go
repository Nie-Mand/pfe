package main

import (
	"Nie-Mand/pfender/cmd"

	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load("config.txt")
	cmd.Execute()

}