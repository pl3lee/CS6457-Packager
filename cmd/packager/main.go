package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pl3lee/CS6457-Packager/internal/utils"
)

func main() {
	config := utils.ParseFlags()
	validator := utils.NewValidator(config)

	missingRequired, foundForbidden, err := validator.ValidateDirectory(config.Directory)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if len(missingRequired) > 0 {
		fmt.Println("Some required files/directories missing:")
		for _, path := range missingRequired {
			fmt.Printf("  %s\n", path)
		}
	}

	if len(foundForbidden) > 0 {
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
			os.Exit(0)
		}
	}

	if len(missingRequired) == 0 && len(foundForbidden) == 0 {
		fmt.Println("Directory looks good!")
	}
}
