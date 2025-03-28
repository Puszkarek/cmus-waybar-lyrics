package models

// LyricsStatus represents the data structure for a lyrics status
type LyricsStatus struct {
	CurrentLine string
	NextLine    string
	TimeToNext  float64
}