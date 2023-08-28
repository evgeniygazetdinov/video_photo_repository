FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add bash && apk add git && mkdir api_calc_go 


COPY . /api_calc_go/

RUN ls

WORKDIR /api_calc_go

RUN ls

RUN go get github.com/gorilla/mux && go get github.com/gorilla/handlers
EXPOSE 8080

CMD go run main.go
