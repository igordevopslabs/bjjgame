package main

import (
	"fmt"

	"github.com/igordevopslabs/bjjgame/pkg"
)

func init() {
	pkg.ConnectToDatabase()
	pkg.MigrateDB()
}

func main() {
	fmt.Println("Jiu Jitsu game!")
}
