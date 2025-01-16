package main

import (
	"log"

	"github.com/golang-documented-todo-api/app/cmd"
)

func main() {
	log.Fatal(cmd.RunServer())
}
