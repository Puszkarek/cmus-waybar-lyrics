package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/puszkarek/cmus-waybar-lyrics/internal/cmus"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/lyrics"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/models"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/utils"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/waybar"
)

func main() {
	// Handle SIGUSR1 for manual updates
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGUSR1)

	// Variables to track the current song
	var currentSongPath string
	var currentSong models.SongMetadata

	// Function to update lyrics
	updateLyrics := func() {
		songStatus, err := cmus.GetSongStatusFromCMUS()
		if err != nil {
			// It's almost sure not running
			waybar.RenderNoSong()
			time.Sleep(time.Second * 2)
			return
		}

		// Check if the song has changed
		if songStatus.Path != currentSongPath {
			// Song has changed, fetch new lyrics
			song, err := lyrics.GetFullLyrics(songStatus.Path)
			if err != nil {
				waybar.RenderNoSong()
				fmt.Println("Error getting lyrics:", err)
				time.Sleep(time.Second * 2)
				return
			}
			
			// Update current song info
			fmt.Println("Song changed, updating waybar...")
			currentSongPath = songStatus.Path
			currentSong = song
		}

		if utils.IsEmptyString(currentSong.Lyrics) && utils.IsNotEmptyString(currentSong.Artist) && utils.IsNotEmptyString(currentSong.Title) {
			waybar.RenderSongStatus(currentSong)
			time.Sleep(time.Second * 2)
			return
		}

		lyricsStatus, err := lyrics.GetCurrentLyricsLine(currentSong.Lyrics, songStatus.Duration, songStatus.Position)
		if err != nil {
			fmt.Println("Error getting current lyrics line:", err)
			time.Sleep(time.Second * 2)
			return
		}

		// Get the time until the next line should be displayed
		waybar.RenderLyrics(lyricsStatus)

		// If we know when the next line should appear, sleep until then
		if lyricsStatus.TimeToNext > 0 {
			sleepTime := math.Max(0.1, lyricsStatus.TimeToNext)
			time.Sleep(time.Duration(sleepTime * float64(time.Second)))
		} else {
			// Otherwise check every second
			time.Sleep(time.Second)
		}
	}

	// Handle SIGUSR1 signal for manual updates
	go func() {
		for range sigChan {
			fmt.Println("Received SIGUSR1, updating waybar...")
			// Force a refresh of lyrics on signal
			currentSongPath = "" // Reset the path to force refresh
			updateLyrics()
		}
	}()

	// Main loop
	for {
		updateLyrics()
	}
}
