package services

import (
	"io"
	"mime/multipart"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func GetFileFromForm(c *gin.Context) (file multipart.File, header *multipart.FileHeader, err error) {
	return c.Request.FormFile("file")
}

func ReadFile(file io.Reader) ([]byte, error) {
	return io.ReadAll(file)
}
