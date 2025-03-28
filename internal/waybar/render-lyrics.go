package waybar

import (
	"encoding/json"
	"fmt"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/utils"
)

// RenderLyrics updates the waybar display with current lyrics
func getDisplayText(lyrics models.LyricsStatus, song models.SongMetadata) models.WaybarOutput {


	if utils.IsEmptyString(lyrics.CurrentLine) {
		return models.WaybarOutput{
			Text:    "...",
			Tooltip: "...",
			Alt: "",
			// At this step we know that the song is playing, so if the current line does not contains anything it
			// means that the song it could be a guitar solo or something like that
			// so we still consider the song as having lyrics
			Class:   "has-lyrics",
		}
	}


	return models.WaybarOutput{
		Text:    lyrics.CurrentLine,
		Tooltip: fmt.Sprintf("%s - %s", song.Artist, song.Title),
		Alt:  lyrics.NextLine,
		Class:   "has-lyrics",
	}

}

// RenderLyrics updates the waybar display with current lyrics
func RenderLyrics(lyrics models.LyricsStatus, song models.SongMetadata) {
	output := getDisplayText(lyrics, song)
	jsonOutput, _ := json.Marshal(output)
	fmt.Println(string(jsonOutput))
}