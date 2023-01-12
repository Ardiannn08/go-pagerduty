FROM golang
ADD . /go/src/github.com/Ardiannn08/go-pagerduty
WORKDIR /go/src/github.com/Ardiannn08/go-pagerduty
RUN go get ./... && go test -v -race -cover ./...
