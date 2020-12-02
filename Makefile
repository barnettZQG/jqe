build:
	GOOS=linux go build -ldflags "-w -s" -o binary/linux-jqe
	GOOS=windows go build -ldflags "-w -s" -o binary/window-jqe.exe
	GOOS=darwin go build -ldflags "-w -s" -o binary/darwin-jqe