package external

// Status represents the external enum type
type Status string

const (
	// StatusActive represents active status
	StatusActive Status = "active"
	// StatusInactive represents inactive status
	StatusInactive Status = "inactive"
	// StatusPending represents pending status
	StatusPending Status = "pending"
)

// Priority represents an external enum for priority levels
type Priority int

const (
	// PriorityLow represents low priority
	PriorityLow Priority = 1
	// PriorityMedium represents medium priority
	PriorityMedium Priority = 2
	// PriorityHigh represents high priority
	PriorityHigh Priority = 3
)
