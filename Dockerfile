FROM golang:1.16-alpine as builder

WORKDIR /app

COPY  go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o fuego-cache -ldflags="-s -w" cmd/main.go

# Stage for the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/fuego-cache /app/config.json /app/

EXPOSE 9919
EXPOSE 9999

CMD /app/fuego-cache