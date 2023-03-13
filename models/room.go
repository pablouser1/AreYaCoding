package models

type Room struct {
	Common
	Description string `json:"description"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
}
