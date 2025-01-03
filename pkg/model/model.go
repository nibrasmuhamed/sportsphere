package models

import "time"

// Turf represents an individual turf under an operator.
type Turf struct {
	ID            string    `json:"id" bson:"_id"`
	OperatorID    string    `json:"operatorId" bson:"operator_id"`
	Name          string    `json:"name" bson:"name"`
	Location      string    `json:"location" bson:"location"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
	IsDeactivated bool      `json:"isDeactivated" bson:"is_deactivated"`
}

// User represents a system user.
type User struct {
	ID            string    `json:"id" bson:"_id"`
	UserName      string    `json:"userName" bson:"user_name"`
	Email         string    `json:"email" bson:"email"`
	Phone         string    `json:"phone" bson:"phone"`
	Password      string    `json:"password" bson:"password"`
	OperatorID    string    `json:"operatorId,omitempty" bson:"operator_id,omitempty"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
	IsDeactivated bool      `json:"isDeactivated" bson:"is_deactivated"`
}

// DaySlots represents the slots for a specific day.
type DaySlots struct {
	ID          string    `json:"id" bson:"_id"`
	Date        string    `json:"date" bson:"date"` // Format: "YYYY-MM-DD"
	TurfID      string    `json:"turfId" bson:"turf_id"`
	BookedSlots []string  `json:"bookedSlots" bson:"booked_slots"`
	CreatedAt   time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updated_at"`
}

// Booking represents a booking made by a user.
type Booking struct {
	ID         string    `json:"id" bson:"_id"`
	UserID     string    `json:"userId" bson:"user_id"`
	TurfID     string    `json:"turfId" bson:"turf_id"`
	OperatorID string    `json:"operatorId" bson:"operator_id"`
	Slots      []string  `json:"slots" bson:"slots"` // List of booked slots
	Date       string    `json:"date" bson:"date"`   // Format: "YYYY-MM-DD"
	CreatedAt  time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updated_at"`
	Duration   int       `json:"duration" bson:"duration"` // Duration in minutes
}

// CacheSlots holds pre-defined slot timings.
type CacheSlots struct {
	ID    string   `json:"id" bson:"_id"`
	Slots []string `json:"slots" bson:"slots"` // Example: ["09:00-10:00", "10:00-11:00"]
}
