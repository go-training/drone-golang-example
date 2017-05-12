
all: docker_build docker_run

test:
	go test -v .

docker_build:
	docker build -t appleboy/go-app .

docker_run:
	docker run --rm appleboy/go-app
