package models

type Room struct {
	Common
	Name string `json:"name"`
	Slug string `json:"slug"`
}
