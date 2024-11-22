package main

import (
	"github.com/gin-gonic/gin"
	"notes_project/controllers"
	"notes_project/database"
	"notes_project/middlewares"
)

func main() {
	// Inisialisasi koneksi ke database
	database.Init()

	// Membuat router menggunakan Gin
	r := gin.Default()

	// Endpoint untuk registrasi dan login
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Endpoint untuk fitur catatan (notes) yang membutuhkan autentikasi
	auth := r.Group("/notes")
	auth.Use(middlewares.AuthMiddleware()) // Middleware untuk validasi token
	{
		auth.POST("/", controllers.CreateNote) // Membuat catatan baru
		auth.GET("/", controllers.GetMyNotes)  // Mendapatkan catatan milik user
	}

	// Menjalankan server pada port 8080
	r.Run(":8080")
}