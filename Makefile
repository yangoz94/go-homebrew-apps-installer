run:
	go run cmd/main.go 

build:
	rm -rf build/
	mkdir build
	GOOS=darwin GOARCH=amd64 go build -o build/ogi-app-amd64 cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o build/ogi-app-arm64 cmd/main.go
	zip -r build/ogi-app.zip build/ogi-app-amd64 build/ogi-app-arm64
