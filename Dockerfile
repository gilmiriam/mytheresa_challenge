FROM golang:1.13.4
WORKDIR /go/src/github.com/mytheresa_challenge
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]