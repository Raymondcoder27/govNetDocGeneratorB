package main

import (
	"example/pdfgenerator/controllers"
	"example/pdfgenerator/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){


	r := gin.Default()


	// Set up CORS middleware
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))



	// r.POST("/generate", controllers.CreateOutput)
	// r.POST("/generate", controllers.CreateOutput)
	r.POST("/generate", controllers.CreateOutput)
	r.GET("/documents", controllers.GetDocuments)
	r.GET("/documents/:id", controllers.GetDocumentById)
	r.DELETE("/documents/:id", controllers.DeleteDocument)
	// r.GET("/generated-file", controllers.GeneratedOutput)
	r.Run()
}
