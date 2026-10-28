package models

import (
	"time"

	"gorm.io/gorm"
)

// ? Note: the foreignKey:X;references:ID tags are optional.
// ?	GORM infers them from the <Name>ID field convention and creates the same foreign keys either way.
type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

type Venue struct {
	gorm.Model
	Name     string `gorm:"index;not null"`
	Location string
	Events   []Event
}

type Performer struct {
	gorm.Model
	Name        string `gorm:"index;not null"`
	Description string
	Events      []Event
}

type Event struct {
	gorm.Model
	Name        string `gorm:"index;not null"`
	Description string
	StartsAt    time.Time
	VenueID     uint
	Venue       Venue
	PerformerID uint
	Performer   Performer
	Tickets     []Ticket
}

type Ticket struct {
	gorm.Model
	EventID    uint `gorm:"index;not null"`
	Event      Event
	PriceCents int
}

type Booking struct {
	gorm.Model
	UserID   uint `gorm:"index;not null"`
	User     User
	TicketID uint `gorm:"uniqueIndex;not null"`
	Ticket   Ticket
}
