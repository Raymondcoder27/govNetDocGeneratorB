package controllers

import (
	// "example/pdfgenerator/services"
	// "example/pdfgenerator/services/decoder"
	// "example/pdfgenerator/services/fileconverter"
	"example/pdfgenerator/initializers"
	"example/pdfgenerator/models"
	"example/pdfgenerator/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PDFResponse struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Data string `json:"data"`
}

func CreateOutput(c *gin.Context) {
	// Log the start of the request
    // fmt.Println("CreateOutput: Received request to generate PDF")

	file, headers, err := services.GetFileFromForm(c)
	if err != nil {
		fmt.Printf("Error retrieving file: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	 // Log file retrieval success
	//  fmt.Println("CreateOutput: File successfully retrieved")

	// Log file details
	// fmt.Printf("CreateOutput: File name is %s\n", header.Filename)

	templateBytes, err := services.ReadFile(file)
	// templateBytes, err := fileconverter.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	jsonData := c.PostForm("data")
	data, err := services.DecodeJSON(jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON data: " + err.Error()})
		return
	}

	pdfBase64, err := services.GeneratePDF(templateBytes, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	 // Generate a new UUID
	 id := uuid.New().String()

	 // Save the PDF to the database
	 pdf := models.PDF{
		 ID:   id,
		 Name: headers.Filename, 
		 Data: pdfBase64,
	 }
	 if err := services.SavePDF(pdf); err != nil {
		 fmt.Printf("Error saving PDF to database: %v\n", err)
		 c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving PDF"})
		 return
	 }

	 // Prepare the response
	 pdfResponse := PDFResponse{
        ID:   id,
        Name: pdf.Name,
        Data: pdfBase64,
    }

	c.JSON(http.StatusOK, gin.H{"pdf": pdfBase64, "id": id})
	fmt.Print("PDF Full Response:", pdfResponse)
}

func GetDocuments(c *gin.Context){
	var documents []models.PDF
	if err := initializers.DB.Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching documents"})
        return
	}
	c.JSON(http.StatusOK, documents)
}

func GetDocumentById(c *gin.Context){
	id := c.Param("id")

	var document models.PDF
	if err := initializers.DB.Where("id = ?", id).First(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching document"})
        return
	}
	//return the doc as a json response
	c.JSON(http.StatusOK, document)
}

func DeleteDocument(c *gin.Context){
	id := c.Param("id")
	var document models.PDF

	// Check if the document exists
	if err := initializers.DB.Where("id = ?", id).First(&document).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
		return
	}

	// Delete the document
	if err := initializers.DB.Delete(&document).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
