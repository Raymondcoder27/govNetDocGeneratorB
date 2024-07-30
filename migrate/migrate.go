package main

import (
	"example/pdfgenerator/initializers"
	"example/pdfgenerator/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.PDF{})
}