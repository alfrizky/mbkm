package models

// Struktur model untuk User
type User struct {
	ID       uint   `gorm:"primaryKey"`     // ID unik untuk setiap user
	Username string `gorm:"unique;not null"` // Username harus unik
	Password string `gorm:"not null"`       // Password yang dienkripsi
	Notes    []Note `gorm:"foreignKey:UserID"` // Relasi dengan tabel Notes
}