package cmus

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
)

// gets song path, duration and position from `cmus-remote`
func GetSongStatusFromCMUS() (models.SongStatus, error) {
	cmd := exec.Command("cmus-remote", "-Q")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return models.SongStatus{}, err
	}

	var songPath string
	var duration, position float64

	status := strings.SplitSeq(string(output), "\n")
	for line := range status {
		lineSplit := strings.SplitN(line, " ", 2)
		if len(lineSplit) < 2 {
			continue
		}

		if lineSplit[0] == "file" {
			songPath = lineSplit[1]
		} else if lineSplit[0] == "duration" {
			durationInt, _ := strconv.Atoi(strings.TrimSpace(lineSplit[1]))
			duration = float64(durationInt)
		} else if lineSplit[0] == "position" {
			positionInt, _ := strconv.Atoi(strings.TrimSpace(lineSplit[1]))
			position = float64(positionInt)
		}
	}

	return models.SongStatus{
		Path: songPath,
		Duration: duration,
		Position: position,
	}, nil
}
