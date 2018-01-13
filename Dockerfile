FROM plugins/base:multiarch

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

ADD release/linux/amd64/helloworld /

ENTRYPOINT ["/helloworld"]
