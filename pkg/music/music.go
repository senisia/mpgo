package music

import (
	"fmt"
	"mpgo/pkg/utils"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlayMusic(name string) {
  musicPath := filepath.Join(utils.GetMusicDir(), name)

	fileHandle, err := os.Open(musicPath)
  if err != nil {
    fmt.Printf("Error with getting file handle: %v", err)
    os.Exit(1)
  }

	streamer, format, err := mp3.Decode(fileHandle)
	if err != nil {
    fmt.Printf("Error: %v", err)
    os.Exit(1)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}
