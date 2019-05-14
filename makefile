build:
	go build -o repohook pkg/main.go

run:
	./repohook -branch=master -path=/Users/prima/Desktop -repo=https://github.com/prima101112/talks -interval=60