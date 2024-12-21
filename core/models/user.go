package models

type Useer struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserName string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Name     string
	Age      uint
}
