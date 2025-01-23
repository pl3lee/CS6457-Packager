.PHONY: build build-macos build-windows build-all

build:
	go build -o packager cmd/packager/main.go

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o packager-intel cmd/packager/main.go
	GOOS=darwin GOARCH=arm64 go build -o packager-arm64 cmd/packager/main.go
	lipo -create -output packager-macos packager-intel packager-arm64
	rm packager-intel packager-arm64

build-windows:
	GOOS=windows GOARCH=amd64 go build -o packager.exe cmd/packager/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o packager-linux cmd/packager/main.go

build-all: build-macos build-windows build-linux