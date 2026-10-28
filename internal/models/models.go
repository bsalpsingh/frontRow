package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"index"`
}

type Venue struct {
	gorm.Model
	Name     string   `gorm:"index"`
	Location string   `gorm:"index"`
	Events   []Events `gorm:"foreignkey:VenueID"`
}

type Performer struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string
	Events      []Events `gorm:"foreignKey:PerformerID"`
}
type Events struct {
	gorm.Model
	Name        string `gorm:"index"`
	Description string `gorm:"index"`
	VenueID     uint
	Venue       Venue `gorm:"foreignKey:VenueID;references:ID"`
	PerformerID uint
	Performer   Performer `gorm:"foreignKey:PerformerID;references:ID"`
}
