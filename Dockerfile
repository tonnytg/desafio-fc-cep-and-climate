FROM golang:1.22.2-alpine as build
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o location main.go

FROM scratch
WORKDIR /app

COPY --from=build /app/location .

EXPOSE 8080

ENTRYPOINT ["./location"]