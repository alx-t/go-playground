TARGET=playground
BETA=beta/playground
BETA_PATH=beta.domido.com

all: clean format build

beta: clean format build_beta

clean:
	rm -rf $(TARGET)

format:
	gofmt -l -s -w .

build:
	go build -o $(TARGET) main.go

clean_beta:
	rm -rf $(BETA)

build_beta:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BETA) main.go

deploy_beta:

run:
	docker run --rm \
	-p 8080:8080 \
	--name=go_playground \
	playground
