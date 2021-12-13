
build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/linux/amd64/helloworld

build_linux_i386:
	go build -v -a -o release/${GOOS}/${GOARCH}/helloworld

build:
	go build -v -a -o release/${GOOS}/${GOARCH}/helloworld

docker:
	docker build -t appleboy/helloworld .

test:
	go test -v .
