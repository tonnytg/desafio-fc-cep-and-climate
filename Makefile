all: build run

build:
	docker build -t cloudrun:latest .

run:
	docker run -d -p 8080:8080 -e WEATHER_API_KEY cloudrun:latest