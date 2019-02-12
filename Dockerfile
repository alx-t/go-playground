FROM golang
MAINTAINER Alex Tesevich <alex.tesevich@gmail.com>
WORKDIR /go/src/github.com/alx-t/go-playground
ADD . ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o playground main.go
EXPOSE 8080 8081
ENTRYPOINT ./playground
