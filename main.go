package main

import (
	"RESTAPICRUD/config"
	"fmt"
)

func main() {
	fmt.Println("Starting application...")
	config.ConnectDb()
}
