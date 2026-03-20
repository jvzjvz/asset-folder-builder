package main

import (
	"fmt"
	"os"
	"strings"
	// "strings"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		return
	}

	entries, err := os.ReadDir(cwd)

	found_asset_folder := false
	for _, entry := range entries {
		fmt.Println(" ", entry.Name(), entry.IsDir())

		if entry.IsDir() && entry.Name() == "assets" {
			found_asset_folder = true
			fmt.Println("BADOOEY!")
		}
	}

	if !found_asset_folder {
		return
	}

	asset_folder_path := fmt.Sprintf("%s/assets/", cwd)
	asset_entries, err := os.ReadDir(asset_folder_path)
	if err != nil {
		return
	}

	fmt.Println("Inside asset folder")
	for _, entry := range asset_entries {
		name := entry.Name()

		if !strings.Contains(name, ".png") {
			continue
		}

		split := strings.Split(name, "-")
		fmt.Println(" ", split)

		directory_names := split[:len(split)-1]
		fmt.Println(" ", directory_names)

	}
}
