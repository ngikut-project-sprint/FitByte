package services

import (
	"fitbyte/internal/database"
	"fitbyte/internal/models"
	"strconv"
)

// GetActivities fetches activities based on filters
func GetActivities(filter models.ActivityFilter) ([]models.Activity, error) {
	query := `SELECT id AS activityId, activity_type AS activityType, done_at AS doneAt, 
          duration_minutes AS durationInMinutes, calories_burned AS caloriesBurned, 
          created_at AS createdAt 
          FROM activities WHERE deleted_at IS NULL`

	args := []interface{}{}
	argPos := 1

	// Dynamic filtering conditions
	if filter.ActivityType != "" {
		query += ` AND activity_type = $` + strconv.Itoa(argPos)
		args = append(args, filter.ActivityType)
		argPos++
	}
	if filter.DoneAtFrom != nil {
		query += ` AND done_at >= $` + strconv.Itoa(argPos)
		args = append(args, *filter.DoneAtFrom)
		argPos++
	}
	if filter.DoneAtTo != nil {
		query += ` AND done_at <= $` + strconv.Itoa(argPos)
		args = append(args, *filter.DoneAtTo)
		argPos++
	}
	if filter.CaloriesBurnedMin != nil {
		query += ` AND calories_burned >= $` + strconv.Itoa(argPos)
		args = append(args, *filter.CaloriesBurnedMin)
		argPos++
	}
	if filter.CaloriesBurnedMax != nil {
		query += ` AND calories_burned <= $` + strconv.Itoa(argPos)
		args = append(args, *filter.CaloriesBurnedMax)
		argPos++
	}

	// Add pagination
	query += ` LIMIT $` + strconv.Itoa(argPos)
	args = append(args, filter.Limit)
	query += ` OFFSET $` + strconv.Itoa(argPos+1)
	args = append(args, filter.Offset)

	// Execute Query
	db := database.GetDB()
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	activities := []models.Activity{}
	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ActivityID, &activity.ActivityType, &activity.DoneAt, &activity.DurationInMinutes, &activity.CaloriesBurned, &activity.CreatedAt); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}