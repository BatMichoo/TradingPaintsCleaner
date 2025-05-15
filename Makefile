APP_NAME=Cleaner

build:
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe
