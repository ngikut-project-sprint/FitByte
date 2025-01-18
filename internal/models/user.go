package models

type User struct {
	Preference string `json:"preference"`
	WeightUnit string `json:"weightUnit"`
	HeightUnit string `json:"heightUnit"`
	Weight     int    `json:"weight"`
	Height     int    `json:"height"`
	Email      string `json:"email"`
	ImageUri   string `json:"imageUri"`
}
