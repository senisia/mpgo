package utils

import (
	"fmt"
	"os"
	"path/filepath"
)


func GetMusicDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error with getting home dir:  %s", err)
	}

	musicDir := filepath.Join(homeDir, "music")

	return musicDir

}


