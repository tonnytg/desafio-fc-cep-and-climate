FROM golang:1.22.2-alpine as build
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o cloudrun cmd/webserv/main.go

FROM scratch
WORKDIR /app

COPY --from=build /app/cloudrun .

EXPOSE 8080

ENTRYPOINT ["./cloudrun"]