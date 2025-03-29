package lyrics

import (
	"errors"
	"math"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
)

// Check if any lines have timestamps
func containsTimestamps(lines []models.LyricLine) bool {
	for _, line := range lines {
		if line.Timestamp >= 0 {
			return true
		}
	}
	return false
}


// Find line info for timestamped lyrics
func findTimestampBasedLineInfo(lines []models.LyricLine, position float64) (int, float64) {
	currentIndex := 0
	timeToNext := math.Inf(1) // Default to infinity if no next line
	
	// Find current line (last line with timestamp <= position)
	for i, line := range lines {
		if line.Timestamp >= 0 && line.Timestamp <= position {
			currentIndex = i
		}
	}
	
	// Find time to next line
	for _, line := range lines {
		if line.Timestamp >= 0 && line.Timestamp > position {
			nextTime := line.Timestamp - position
			if nextTime < timeToNext {
				timeToNext = nextTime
			}
		}
	}
	
	return currentIndex, timeToNext
}

// Find line info for untimed lyrics
func findUDurationBasedLineInfo(lines []models.LyricLine, position, duration float64) (int, float64) {
	lineCount := len(lines)
	currentIndex := int(position * float64(lineCount) / duration)
	
	// Calculate time to next line
	timeToNext := math.Inf(1) // Default to infinity
	if currentIndex < lineCount-1 {
		timePerLine := duration / float64(lineCount)
		timeToNext = timePerLine - math.Mod(position, timePerLine)
	}
	
	return currentIndex, timeToNext
}

// Get next line text safely
func getNextLineText(lines []models.LyricLine, currentIndex int) string {
	if currentIndex+1 < len(lines) {
		return lines[currentIndex+1].Text
	}
	return ""
}



// Find current line index and time to next line
func findCurrentLineInfo(lines []models.LyricLine, position, duration float64, hasTimestamps bool) (int, float64) {
	if hasTimestamps {
		return findTimestampBasedLineInfo(lines, position)
	}
	return findUDurationBasedLineInfo(lines, position, duration)
}


// Gets the current line of lyrics based on song position using immutable data approach
func GetDisplayLyrics(lyricLines []models.LyricLine, songDuration, songPosition float64) (models.LyricsStatus, error) {
	if len(lyricLines) == 0 || songDuration <= 0 {
		return models.LyricsStatus{}, errors.New("invalid lyrics or song duration")
	}

	// Determine if we have timestamped lyrics
	hasTimestamps := containsTimestamps(lyricLines)
	
	// Find current line and time to next line
	currentIndex, timeToNext := findCurrentLineInfo(lyricLines, songPosition, songDuration, hasTimestamps)
	
	if currentIndex < 0 || currentIndex >= len(lyricLines) {
		return models.LyricsStatus{}, errors.New("no current line found")
	}
	
	// Build result with immutable values
	return models.LyricsStatus{
		CurrentLine: lyricLines[currentIndex].Text,
		NextLine:    getNextLineText(lyricLines, currentIndex),
		TimeToNext:  timeToNext,
	}, nil
}