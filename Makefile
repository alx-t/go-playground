TARGET=playground

all: clean format build

clean:
	rm -rf $(TARGET)

format:
	gofmt -l -s -w .

build:
	go build -o $(TARGET) main.go
