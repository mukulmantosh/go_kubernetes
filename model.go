package main

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Age   int    `json:"age"`
}
