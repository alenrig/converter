EXECUTABLE=converter

all:
	go build -o $(EXECUTABLE) main.go

clean:
	rm -f $(EXECUTABLE)
