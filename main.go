package main

import (
	"log"

	//"github.com/Faheem-Nizar/go-rabbitmq-tutorial/internal/cmd"
	"proton/internal/cmd"
)


func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	cmd.Execute()
}