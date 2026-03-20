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

	os.Chdir(asset_folder_path)

	fmt.Println("Inside asset folder")
	for _, entry := range asset_entries {
		name := entry.Name()

		if !strings.Contains(name, ".png") {
			continue
		}

		last_separator_index := strings.LastIndex(name, "-")
		asset_name := name[last_separator_index+1:]

		directory_name := strings.ReplaceAll(name[:last_separator_index], "-", "/")
		fmt.Println("directory", directory_name)

		fmt.Printf("%s -> %s + %s\n", name, directory_name, asset_name)

		err := os.MkdirAll(directory_name, os.ModePerm)
		if err != nil {
			fmt.Println("directory already exists")
		} else {
			fmt.Println("sucessfully created directory")
		}

		new_asset_location := directory_name + "/" + asset_name
		err = os.Rename(name, new_asset_location)
		if err != nil {
			continue
		} else {
			fmt.Printf("moved %s to %s", name, new_asset_location)
		}
	}
}
