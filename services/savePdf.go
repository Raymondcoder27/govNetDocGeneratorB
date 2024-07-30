package services

import (
	"example/pdfgenerator/initializers"
	"example/pdfgenerator/models"

	"gorm.io/gorm"
)

var DB *gorm.DB // Assume this is initialized somewhere

func SavePDF(pdf models.PDF) error {
    return initializers.DB.Create(&pdf).Error
}
