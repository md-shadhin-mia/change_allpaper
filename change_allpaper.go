package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Function to get all wallpapers from the folder
func getWallpapers(folder string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(folder, "*.jpg"))
	if err != nil {
		return nil, err
	}

	pngFiles, err := filepath.Glob(filepath.Join(folder, "*.png"))
	if err != nil {
		return nil, err
	}

	files = append(files, pngFiles...)

	if len(files) == 0 {
		return nil, fmt.Errorf("no wallpapers found in the folder")
	}

	return files, nil
}

// Function to set the wallpaper using gsettings
func setWallpaper(wallpaperPath string) error {
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fmt.Sprintf("file://%s", wallpaperPath))
	return cmd.Run()
}

func main() {
	// Define command-line flags
	fmt.Println("FREE PALESTINE 🇵🇸")
	duration := flag.Int("duration", 5, "Duration in seconds between wallpaper changes")
	path := flag.String("path", ".", "Path to the wallpapers folder")

	flag.Parse()

	interval := time.Duration(*duration) * time.Second
	wallpapers, err := getWallpapers(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	for {
		for _, wallpaper := range wallpapers {
			if err := setWallpaper(wallpaper); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				return
			}
			time.Sleep(interval)
		}
	}
}
