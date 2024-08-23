package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // POST endpoint to handle file uploads
    r.POST("/uploads", func(c *gin.Context) {
        // Get the file from the form input
        file, err := c.FormFile("file")
        if err != nil {
            c.String(http.StatusBadRequest, "Failed to get file: %s", err.Error())
            return
        }

        // Define the path where the file will be saved
        uploadPath := "./user_uploads"
        if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
            os.Mkdir(uploadPath, os.ModePerm)
        }

        // Construct the full path for the uploaded file
        fullPath := filepath.Join(uploadPath, file.Filename)

        // Save the file to the specified path
        if err := c.SaveUploadedFile(file, fullPath); err != nil {
            c.String(http.StatusInternalServerError, "Failed to save file: %s", err.Error())
            return
        }

        c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded successfully!", file.Filename))
    })

    r.StaticFS("files", http.Dir("./user_uploads"))

    r.Run(":2100")
}
