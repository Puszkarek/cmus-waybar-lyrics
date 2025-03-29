package models

// SongMetadata represents the data structure for a song
type SongMetadata struct {
	ParsedLyrics []LyricLine
	Title       string
	Artist      string
}