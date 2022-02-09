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

compose-up:
	docker-compose -f docker-compose-local.yaml up

compose:
	docker-compose build
	docker-compose up
	docker image prune -f

docker-clean:
	#docker rm  $(docker ps -a -s | grep -v "mongo" | awk '{print $1}' | grep -v CONTAINER)
	docker image prune -f

