package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pl3lee/CS6457-Packager/internal/config"
)

type Validator struct {
	Directory string
	Name      string
	Mode      string
}

func NewValidator(config config.AppConfig) *Validator {
	return &Validator{
		Directory: config.Directory,
		Name:      config.Name,
		Mode:      config.Mode,
	}
}

// puts forbidden paths/files into foundForbidden
func (v *Validator) ValidateDirectory(dir string) (missingRequired []string, foundForbidden []string, err error) {
	// Check required paths
	requiredPaths := config.GetAllRequiredPaths(v.Name)
	for _, path := range requiredPaths {
		fullPath := filepath.Join(dir, path)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			missingRequired = append(missingRequired, path)
		}
	}

	// Check forbidden paths
	for _, forbidden := range config.ForbiddenPaths {
		fullPath := filepath.Join(dir, forbidden)
		if _, err := os.Stat(fullPath); err == nil {
			foundForbidden = append(foundForbidden, forbidden)
		}
	}

	// Check forbidden extensions
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		for _, forbidden := range config.ForbiddenExtensions {
			if ext == forbidden {
				relPath, _ := filepath.Rel(dir, path)
				foundForbidden = append(foundForbidden, relPath)
			}
		}
		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("error walking directory: %v", err)
	}

	return missingRequired, foundForbidden, nil
}

// Removes all paths from path slice.
func (v *Validator) CleanDirectory(paths []string) error {
	for _, path := range paths {
		fullPath := filepath.Join(v.Directory, path)
		if err := os.RemoveAll(fullPath); err != nil {
			return fmt.Errorf("failed to remove %s: %v", fullPath, err)
		}
	}
	return nil
}
