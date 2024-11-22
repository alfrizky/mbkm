package database

import (
)

var DB *gorm.DB

// Fungsi untuk menginisialisasi koneksi ke database
func Init() {
	var err error
	// Menggunakan SQLite sebagai database
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database!")
	}

	// Migrasi tabel berdasarkan model User dan Note
	DB.AutoMigrate(&models.User{}, &models.Note{})
}
