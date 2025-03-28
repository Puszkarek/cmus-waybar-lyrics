package waybar

import (
	"encoding/json"
	"fmt"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
)

func RenderNoSong()  {
	output := models.WaybarOutput{
		Text:    "No song playing",
		Tooltip: "No song playing",
		Class:   "no-song",
	}
	jsonOutput, _ := json.Marshal(output)
	fmt.Println(string(jsonOutput))
}