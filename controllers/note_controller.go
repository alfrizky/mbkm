package controllers

import (
	"github.com/gin-gonic/gin"
	"notes_project/database"
	"notes_project/models"
	"net/http"
)

// Endpoint untuk membuat catatan baru
func CreateNote(c *gin.Context) {
	var note models.Note
	userID := c.GetUint("userID") // Mendapatkan ID user dari middleware
	note.UserID = userID

	// Parsing data dari body request ke struct Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menyimpan catatan ke database
	if err := database.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// Endpoint untuk mendapatkan semua catatan milik user
func GetMyNotes(c *gin.Context) {
	var notes []models.Note
	userID := c.GetUint("userID") // Mendapatkan ID user dari middleware

	// Mengambil catatan dari database berdasarkan ID user
	if err := database.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}
