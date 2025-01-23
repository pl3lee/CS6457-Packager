# CS6457 Unity Project Packager

A command-line tool to validate and clean Unity project structures for CS6457 assignments. Ensures project structure meets submission requirements and removes unnecessary files.

[Assignment Packaging and Submission Instructions Here](https://gatech.instructure.com/courses/436374/pages/assignment-packaging-and-submission)

## Disclaimer

**⚠️ IMPORTANT: Read Before Using**
- Create a backup of your project before using this tool
- Test in a safe environment first
- Verify results manually after using
- The author assumes no responsibility for any data loss or project issues

## Installation

### Download Pre-built Binaries
1. Go to [Releases](https://github.com/pl3lee/CS6457-Packager/releases)
2. Download the appropriate binary for your platform:
   - Windows: `packager.exe`
   - macOS: `packager-macos` (Universal binary for Intel/Apple Silicon)
   - Linux: `packager-linux`

### Build from Source
You need to first install Go on your machine. Follow the instructions [here](https://go.dev/doc/install).
```bash
# Clone repository
git clone https://github.com/pl3lee/CS6457-Packager
cd CS6457-Packager

# Build for linux
make build-linux

# Build for Windows
make build-windows

# Build for MacOS (needs to be on MacOS to work)
make build-macos
```

## Modes
- **Check**: Validates project structure and reports issues
- **Clean**: Removes unnecessary files and folders

See `internal/config/config.go` for the list of files and folders that are removed in clean mode.

## Usage Examples
### Basic Command Structure
```bash
./packager -dir <directory> -mode <check|clean> -name <project_name>
```

In the examples below, replace `./Lee_P_m1` with the path to your Unity project directory, and `Lee_P_m1` with the name of your project.
### Windows
```cmd
.\packager.exe -dir .\Lee_P_m1 -mode check -name Lee_P_m1
.\packager.exe -dir .\Lee_P_m1 -mode clean -name Lee_P_m1
```

### MacOS
```bash
# Make binary executable
chmod +x ./packager-macos
./packager-macos -dir ./Lee_P_m1 -mode check -name Lee_P_m1
./packager-macos -dir ./Lee_P_m1 -mode clean -name Lee_P_m1
```
You may receive a warning saying that the binary cannot be trusted...

To bypass this, go to the "Privacy and Security" settings in the System Prefereces/Settings app, and click "Open Anyway" next to the warning message.

### Linux
```bash
# Make binary executable
chmod +x ./packager-linux
./packager-linux -dir ./Lee_P_m1 -mode check -name Lee_P_m1
./packager-linux -dir ./Lee_P_m1 -mode clean -name Lee_P_m1
```

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

