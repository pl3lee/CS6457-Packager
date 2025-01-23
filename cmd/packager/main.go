package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/pl3lee/CS6457-Packager/internal/utils"
)

func main() {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("DISCLAIMER:")
	fmt.Println("This tool is provided as-is without any warranty.")
	fmt.Println("- Create a backup of your project before using this tool")
	fmt.Println("- Test thoroughly in a safe environment first")
	fmt.Println("- Verify results manually")
	fmt.Println("The author assumes no responsibility for any data loss or project issues.")
	fmt.Println(strings.Repeat("=", 80))
	config := utils.ParseFlags()
	dirName := filepath.Base(config.Directory)
	if dirName != config.Name {
		fmt.Printf("Warning: Directory name '%s' does not match project name '%s', you might want to change that...\n", dirName, config.Name)
	}
	validator := utils.NewValidator(config)

	missingRequired, foundForbidden, err := validator.ValidateDirectory(config.Directory)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if len(foundForbidden) > 0 {
		fmt.Println("\n" + strings.Repeat("=", 80))
		fmt.Println("Found forbidden files/directories:")
		for _, path := range foundForbidden {
			fmt.Printf("  %s\n", path)
		}

		if config.Mode == "clean" {
			if err := validator.CleanDirectory(foundForbidden); err != nil {
				log.Fatalf("Error cleaning: %v\n", err)
			}
			fmt.Println("Directory cleaned successfully")
		} else {
			fmt.Println("Run with mode clean to clean files")
		}
	}

	if len(missingRequired) > 0 {
		fmt.Println("\n" + strings.Repeat("=", 80))
		fmt.Println("Some required files/directories missing:")
		for _, path := range missingRequired {
			fmt.Printf("  %s\n", path)
		}
	}

	if len(missingRequired) == 0 && len(foundForbidden) == 0 {
		fmt.Println("\n" + strings.Repeat("=", 80))
		fmt.Println("Directory looks good!")
		fmt.Println(strings.Repeat("=", 80))
	}
}
