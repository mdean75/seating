[dep-check:
	go list -u -m all

run:
	go run cmd/main.go

build: 
	go build -o out/bin/seating cmd/main.go

clean:
	rm seating 

docker-build:
	docker build . -t seating:latest -f Dockerfile

compose:
	docker-compose build
	docker-compose up
	docker image prune -f
