package lyrics

import (
	"path/filepath"
	"strings"
)

// GetSongMetadataFromPath tries to get song artist and title from its path
func GetSongMetadataFromPath(path string) (string, string) {
	songName := strings.TrimPrefix(filepath.Base(path), "/")
	ext := filepath.Ext(songName)
	songName = strings.TrimSuffix(songName, ext)

	songNameSplit := strings.Split(songName, " - ")
	if len(songNameSplit) < 2 {
		songNameSplit = strings.Split(songName, "-")
	}

	if len(songNameSplit) >= 2 {
		return songNameSplit[0], songNameSplit[1]
	}

	// If we can't split by dash, try to use parent directory as artist
	dir := filepath.Dir(path)
	artist := filepath.Base(dir)
	return artist, songName
}

