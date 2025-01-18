package models

import (
	"time"
)

// ActivityFilter holds filters for activities
type ActivityFilter struct {
	Limit             int
	Offset            int
	ActivityType      string
	DoneAtFrom        *time.Time
	DoneAtTo          *time.Time
	CaloriesBurnedMin *int
	CaloriesBurnedMax *int
}

// Activity represents the structure of an activity
type Activity struct {
	ActivityID       string    `json:"activityId"`
	ActivityType     string    `json:"activityType"`
	DoneAt           time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned   float64   `json:"caloriesBurned"`
	CreatedAt        time.Time `json:"createdAt"`
}