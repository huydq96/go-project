FROM golang:alpine as builder

WORKDIR /
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /cmd/server /cmd/server.go

FROM alpine
WORKDIR /app
COPY --from=builder /cmd/server /app/server
COPY --from=builder /config /app/config
CMD ["/app/server"]