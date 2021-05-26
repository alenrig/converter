EXECUTABLE=dist/converter

all:
	go build -o $(EXECUTABLE) main.go

format:
	go fmt $(go list ./... | grep -v /vendor/) && go vet $(go list ./... | grep -v /vendor/)

clean:
	go clean
