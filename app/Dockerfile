FROM golang:1.12.5

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -v github.com/rubenv/sql-migrate/...

ENV TZ=Asia/Tokyo

WORKDIR $GOPATH/src/github.com/projects/app
