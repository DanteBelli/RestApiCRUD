package main

import (
	"RESTAPICRUD/config"
	"RESTAPICRUD/routes"
	"fmt"
)

func main() {
	fmt.Println("Starting application...")
	config.ConnectDb()
	r := routes.SetupRoute()
	r.Run(":8080")
}
