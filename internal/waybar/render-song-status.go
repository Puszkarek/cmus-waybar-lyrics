package waybar

import (
	"encoding/json"
	"fmt"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
)

// RenderSongStatus updates the waybar display with current lyrics
func RenderSongStatus(song models.SongMetadata)  {

		output := models.WaybarOutput{
			Text:    fmt.Sprintf("%s - %s", song.Artist, song.Title),
			Tooltip: fmt.Sprintf("%s - %s", song.Artist, song.Title),
			Class:   "song",
		}
		jsonOutput, _ := json.Marshal(output)
		fmt.Println(string(jsonOutput))


}