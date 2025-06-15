package models

type User struct {
	CitizenID   string `gorm:"column:CITIZEN_ID;primaryKey;size:13"`
	Name        string `gorm:"column:NAME"`
	LastName    string `gorm:"column:LAST_NAME"`
	Address     string `gorm:"column:ADDRESS"`
	ActiveBikes int    `gorm:"column:ACTIVE_BIKES;default:0"`
}

func (User) TableName() string {
	return "Users"
}
