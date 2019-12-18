FROM golang:1.13.4
WORKDIR /go/src/github.com/tangelo_challenge
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]