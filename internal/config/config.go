package config

import "fmt"

type AppConfig struct {
	Directory string
	Mode      string
	Name      string
}

func GetAllRequiredPaths(name string) []string {
	result := make([]string, len(RequiredPaths))
	copy(result, RequiredPaths)

	for _, template := range DynamicPaths {
		result = append(result, fmt.Sprintf(template, name))
	}
	return result
}

var (
	ForbiddenPaths = []string{
		"Library",
		"Temp",
		"Obj",
		"Logs",
		"UserSettings",
		"MemoryCaptures",
		"Recordings",
		".vs",
		".gradle",
		"ExportedObj",
		".consulo",
		"crashlytics-build.properties",
		"sysinfo.txt",
	}

	// ForbiddenExtensions for files
	ForbiddenExtensions = []string{
		".csproj",
		".unityproj",
		".sln",
		".suo",
		".tmp",
		".user",
		".userprefs",
		".pidb",
		".booproj",
		".svd",
		".pdb",
		".mdb",
		".opendb",
		".VC.db",
		".apk",
		".aab",
		".unitypackage",
		".unitypackage.meta",
		".app",
		".pidb.meta",
		".pdb.meta",
		".mdb.meta",
	}

	// RequiredPaths that must exist
	RequiredPaths = []string{
		"Build/OSX",
		"Build/Windows",
		"Assets",
		"ProjectSettings",
		"Packages",
	}
	// used to generate specific paths
	DynamicPaths = []string{
		"%s_readme.txt",
		"Build/OSX/%s.app",
		"Build/Windows/%s.exe",
	}
)
