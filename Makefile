build:
	GOOS=windows GOARCH=amd64 go build -o bin/mesmerizer-windows64.exe
	GOOS=linux GOARCH=amd64 go build -o bin/mesmerizer-linux64
	GOOS=darwin GOARCH=amd64 go build -o bin/mesmerizer-macos64
