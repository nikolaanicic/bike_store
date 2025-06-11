package models

type User struct {
	CitizenID    string `gorm:"column:CITIZEN_ID;primaryKey;size:13"`
	Name         string `gorm:"column:NAME"`
	LastName     string `gorm:"column:LAST_NAME"`
	Address      string `gorm:"column:ADDRESS"`
	PasswordHash string `gorm:"column:PASSWORD_HASH;size:64"`
}
