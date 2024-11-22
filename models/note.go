package models

// Struktur model untuk Catatan (Note)
type Note struct {
	ID     uint   `gorm:"primaryKey"`  // ID unik untuk setiap catatan
	Title  string `gorm:"not null"`    // Judul catatan
	Body   string                      // Isi catatan
	UserID uint   `gorm:"not null"`    // ID user yang memiliki catatan ini
}
