[dep-check:
	go list -u -m all

run:
	go run seating.go

build: 
	go build -o seating seating.go

clean:
	rm seating 

docker-build:
	docker build . -t seating:latest -f Dockerfile

compose:
	docker-compose build
	docker-compose up