FROM plugins/base:multiarch

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Workshop" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

ADD release/linux/amd64/helloworld /bin/

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/helloworld", "-ping" ]

ENTRYPOINT ["/bin/helloworld"]
