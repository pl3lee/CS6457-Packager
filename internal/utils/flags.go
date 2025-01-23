package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/pl3lee/CS6457-Packager/internal/config"
)

func ParseFlags() config.AppConfig {
	directory := flag.String("directory", "", "Directory path to process")
	mode := flag.String("mode", "", "Mode: check or clean")
	name := flag.String("name", "", "Project name (individual milestones: LASTNAME_FIRSTINITIAL_m[0-9]) (team deliverables: TEAMNAME_GAMENAME)")
	flag.Parse()

	if *directory == "" {
		fmt.Println("Error: directory path is required")
		flag.Usage()
		os.Exit(1)
	}

	if *mode != "check" && *mode != "clean" {
		fmt.Println("Error: mode must be either 'check' or 'clean'")
		flag.Usage()
		os.Exit(1)
	}

	if *name == "" {
		fmt.Println("Please give a project name")
		flag.Usage()
		os.Exit(1)
	}

	return config.AppConfig{
		Directory: *directory,
		Mode:      *mode,
		Name:      *name,
	}
}
