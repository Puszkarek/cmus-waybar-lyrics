package lyrics

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
)

// gets the current line of lyrics based on song position
func GetCurrentLyricsLine(lyrics string, songDuration, songPosition float64) (models.LyricsStatus, error) {
	if lyrics == "" || songDuration <= 0 {
		return models.LyricsStatus{}, errors.New("invalid lyrics or song duration")
	}

	// Regex pattern to match timestamp formats like [02:48.93]
	timestampPattern := regexp.MustCompile(`\[(\d+):(\d+)\.(\d+)\]`)

	// Split lyrics into lines
	lines := strings.Split(lyrics, "\n")
	if len(lines) == 0 {
		return models.LyricsStatus{}, errors.New("no lyrics found")
	}

	// Extract timestamps from each line if available
	timestamps := make([]float64, len(lines))
	hasTimestamps := false

	for i, line := range lines {
		timestampMatch := timestampPattern.FindStringSubmatch(line)
		if len(timestampMatch) > 0 {
			minutes, _ := strconv.Atoi(timestampMatch[1])
			seconds, _ := strconv.Atoi(timestampMatch[2])
			milliseconds, _ := strconv.Atoi(timestampMatch[3])
			timeInSeconds := float64(minutes)*60 + float64(seconds) + float64(milliseconds)/100
			timestamps[i] = timeInSeconds
			hasTimestamps = true
		} else {
			timestamps[i] = -1 // Using -1 to indicate no timestamp
		}
	}

	// Determine current line index
	currentLineIdx := 0
	var timeToNextLine *float64

	if hasTimestamps {
		// Find the appropriate line based on current position
		nextLineFound := false
		for i, timestamp := range timestamps {
			if timestamp >= 0 && timestamp <= songPosition {
				currentLineIdx = i
			} else if timestamp >= 0 && timestamp > songPosition && !nextLineFound {
				nextTime := timestamp - songPosition
				timeToNextLine = &nextTime
				nextLineFound = true
			}
		}
	} else {
		// Fall back to the original method if no timestamps
		currentLineIdx = int(songPosition * float64(len(lines)) / songDuration)
		// Estimate time to next line
		if currentLineIdx < len(lines)-1 {
			timePerLine := songDuration / float64(len(lines))
			nextTime := timePerLine - math.Mod(songPosition, timePerLine)
			timeToNextLine = &nextTime
		}
	}

	// Return the current line if it's valid
	if currentLineIdx >= 0 && currentLineIdx < len(lines) {
		// Clean the line by removing timestamps
		cleanLine := timestampPattern.ReplaceAllString(lines[currentLineIdx], "")
		
		cleanNextLine := ""
    if currentLineIdx+1 < len(lines) {
        cleanNextLine = timestampPattern.ReplaceAllString(lines[currentLineIdx+1], "")
    }
		return models.LyricsStatus{
			CurrentLine: strings.TrimSpace(cleanLine),
			NextLine: strings.TrimSpace(cleanNextLine),
			TimeToNext:  *timeToNextLine,
		}, nil
	}

	return models.LyricsStatus{}, errors.New("no current line found")
}

