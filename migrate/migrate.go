package main

import (
	"miboutaj.com/go-crud/initializers"
	"miboutaj.com/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Student{})
	if err != nil {
		return
	}
}
