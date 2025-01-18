package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"fitbyte/internal/models"
	"fitbyte/internal/services"
)

// ActivityHandler handles /v1/activity endpoints
func ActivityHandler(c echo.Context) error {
	// Parse query parameters
	limit, offset := parsePagination(c.QueryParam("limit"), c.QueryParam("offset"))
	activityType := c.QueryParam("activityType")
	doneAtFrom, _ := parseISODate(c.QueryParam("doneAtFrom"))
	doneAtTo, _ := parseISODate(c.QueryParam("doneAtTo"))
	caloriesBurnedMin := parseOptionalInt(c.QueryParam("caloriesBurnedMin"))
	caloriesBurnedMax := parseOptionalInt(c.QueryParam("caloriesBurnedMax"))

	// Build the filter options
	filterOptions := models.ActivityFilter{
		Limit:             limit,
		Offset:            offset,
		ActivityType:      activityType,
		DoneAtFrom:        doneAtFrom,
		DoneAtTo:          doneAtTo,
		CaloriesBurnedMin: caloriesBurnedMin,
		CaloriesBurnedMax: caloriesBurnedMax,
	}

	// Call service to fetch activities
	activities, err := services.GetActivities(filterOptions)

	if activityType != "" && !isValidActivityType(activityType) {
		activityType = "" // Ignore invalid value
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}

	return c.JSON(http.StatusOK, activities)
}

func parsePagination(limitStr, offsetStr string) (int, int) {
	limit := 5 // Default
	offset := 0
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
		offset = o
	}
	return limit, offset
}

func parseISODate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	parsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func parseOptionalInt(value string) *int {
	if value == "" {
		return nil
	}
	if val, err := strconv.Atoi(value); err == nil {
		return &val
	}
	return nil
}

var validActivityTypes = map[string]bool{
    "WALKING":   true,
    "YOGA":      true,
    "STRECHING": true,
    "CYCLING":   true,
    "SWIMMING":  true,
    "DANCING":   true,
    "HIKING":    true,
    "RUNNING":   true,
    "HIIT":      true,
    "JUMP_ROPE": true,
}

func isValidActivityType(activityType string) bool {
    return validActivityTypes[activityType]
}