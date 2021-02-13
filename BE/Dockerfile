FROM golang

ENV GO111MODULE=on
ENV GIN_MODE=release

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/start.go

EXPOSE 8004
ENTRYPOINT [ "/app/start" ]