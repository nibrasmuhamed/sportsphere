package models

import "time"

// AuditLog represents a record of an action in the system.
// TODO: This feature will not be available in the initial release.
type AuditLog struct {
	ID            string      `json:"id" bson:"_id"`                                           // Unique ID for the audit log
	Timestamp     time.Time   `json:"timestamp" bson:"timestamp"`                              // Time of the action
	Action        string      `json:"action" bson:"action"`                                    // Action type (e.g., "BookingCreated", "SlotUpdated")
	UserID        string      `json:"userId" bson:"user_id"`                                   // User who performed the action
	OperatorID    string      `json:"operatorId" bson:"operator_id"`                           // Operator ID (if applicable)
	TurfID        string      `json:"turfId,omitempty" bson:"turf_id,omitempty"`               // Turf ID (if applicable)
	Entity        string      `json:"entity" bson:"entity"`                                    // Entity affected (e.g., "Booking", "DaySlots")
	EntityID      string      `json:"entityId" bson:"entity_id"`                               // ID of the affected entity
	Changes       interface{} `json:"changes" bson:"changes"`                                  // Detailed changes (optional, can be a JSON object)
	PreviousState interface{} `json:"previousState,omitempty" bson:"previous_state,omitempty"` // State before the action
	IPAddress     string      `json:"ipAddress,omitempty" bson:"ip_address,omitempty"`         // User's IP address (optional)
	UserAgent     string      `json:"userAgent,omitempty" bson:"user_agent,omitempty"`         // User's device or client details (optional)
}
