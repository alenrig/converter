EXECUTABLE=dist/converter

all:
	go build -o $(EXECUTABLE) main.go

clean:
	go clean
