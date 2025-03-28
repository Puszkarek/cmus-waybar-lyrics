package models

// WaybarOutput represents the JSON structure for Waybar
type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Alt     string `json:"alt"`
	Class   string `json:"class"`
}