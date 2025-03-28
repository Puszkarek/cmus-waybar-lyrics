package lyrics

import (
	"errors"
	"os"

	"github.com/dhowden/tag"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/utils"
)

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
				Lyrics: "No lyrics tag.",
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

	return models.SongMetadata{
		Lyrics: lyrics,
		Title:  title,
		Artist: artist,
	}, nil
	}
