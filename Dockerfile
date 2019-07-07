# builder
FROM golang:1.12 as builder

RUN mkdir /osrscx
WORKDIR /osrscx

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/osrscx

# runner
FROM alpine:latest

LABEL maintainer="Corey (notmeta) <https://github.com/notmeta/>"

RUN apk --no-cache add ca-certificates bash
COPY --from=builder /go/bin/osrscx /go/bin/osrscx

ENTRYPOINT ["/go/bin/osrscx"]
