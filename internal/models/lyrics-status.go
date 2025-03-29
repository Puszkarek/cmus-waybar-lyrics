package models

// LyricsStatus represents the data structure for a lyrics status
type LyricsStatus struct {
	// TODO: PreviousLine string
	CurrentLine  string
	NextLine     string
	TimeToNext  float64
}