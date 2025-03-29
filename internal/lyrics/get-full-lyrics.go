package lyrics

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dhowden/tag"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/utils"
)

// Parse raw lines into immutable LyricLine structs
func parseLyricLines(lines []string) []models.LyricLine {
	timestampPattern := regexp.MustCompile(`\[(\d+):(\d+)\.(\d+)\]`)
	result := make([]models.LyricLine, len(lines))
	
	for i, line := range lines {
		timestamp := -1.0
		timestampMatch := timestampPattern.FindStringSubmatch(line)
		
		if len(timestampMatch) > 0 {
			minutes, _ := strconv.Atoi(timestampMatch[1])
			seconds, _ := strconv.Atoi(timestampMatch[2])
			milliseconds, _ := strconv.Atoi(timestampMatch[3])
			timestamp = float64(minutes)*60 + float64(seconds) + float64(milliseconds)/100
		}
		
		cleanText := strings.TrimSpace(timestampPattern.ReplaceAllString(line, ""))
		
		result[i] = models.LyricLine{
			Text:      cleanText,
			Timestamp: timestamp,
		}
	}
	
	return result
}


// GetFullLyrics tries to get song lyrics from tags
func GetFullLyrics(songPath string) (models.SongMetadata, error) {

	// Open the file
	file, err := os.Open(songPath)
	if err != nil {
		return models.SongMetadata{}, err
	}
	defer file.Close()

	// Read the metadata
	metadata, err := tag.ReadFrom(file)
	if err != nil {
		// If there's no metadata, try to get artist and title from path
		 pathArtist, pathTitle := GetSongMetadataFromPath(songPath)

		 // if found return it
		if utils.IsNotEmptyString(pathArtist) && utils.IsNotEmptyString(pathTitle) {
			return models.SongMetadata{
				ParsedLyrics: []models.LyricLine{},
				Title:  pathTitle,
				Artist: pathArtist,
			}, nil
		}
	}

		// Try to get lyrics from metadata

		lyrics := metadata.Lyrics()
		artist := metadata.Artist()
		title := metadata.Title()

		if utils.IsEmptyString(lyrics) && utils.IsEmptyString(artist) && utils.IsEmptyString(title) {
			return models.SongMetadata{}, errors.New("no lyrics, artist or title found in tags")
		}


	// Split lyrics into lines
	lines := strings.Split(lyrics, "\n")
	parsedLyrics := parseLyricLines(lines)

	return models.SongMetadata{
		ParsedLyrics: parsedLyrics,
		Title:  title,
		Artist: artist,
	}, nil
	}
