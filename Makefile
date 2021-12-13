build:
	go build -v -a -o release/${GOOS}/${GOARCH}/helloworld

docker:
	docker build -t appleboy/helloworld .

test:
	go test -v .
