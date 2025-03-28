package models

// WaybarOutput represents the JSON structure for Waybar
type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}