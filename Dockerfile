FROM alpine:latest

COPY bees /usr/bin/bees
RUN chmod +x /usr/bin/bees

ENTRYPOINT ["/usr/bin/bees"]
