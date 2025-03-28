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
	"github.com/puszkarek/cmus-waybar-lyrics/internal/utils"
	"github.com/puszkarek/cmus-waybar-lyrics/internal/waybar"
)

func update() {
		songStatus, err := cmus.GetSongStatusFromCMUS()
		if err != nil {
			// It's almost sure not running
			waybar.RenderNoSong()
			time.Sleep(time.Second *2)
			return
		}
		song, err := lyrics.GetFullLyrics(songStatus.Path)
		if err != nil {
			waybar.RenderNoSong()
			fmt.Println("Error getting lyrics:", err)
			time.Sleep(time.Second * 2)
			return
		}

		if utils.IsEmptyString(song.Lyrics) && utils.IsNotEmptyString(song.Artist) && utils.IsNotEmptyString(song.Title) {
			waybar.RenderSongStatus(song)
			return
		}

		lyricsStatus, err := lyrics.GetCurrentLyricsLine(song.Lyrics, songStatus.Duration, songStatus.Position)
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



func main() {
		// Handle SIGUSR1 for manual updates
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGUSR1)

		go func() {
			for range sigChan {
				fmt.Println("Received SIGUSR1, updating waybar...")
				update()
			}
		}()

// TODO: Find a way to only update the song metadata when the song changes
// TODO: Find a way to only update the lyrics when the song changes
// TODO: Improve synchronization between song and lyrics updates

		for {
		update()			
		}

}
