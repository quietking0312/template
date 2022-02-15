FROM alpine:latest

RUN mkdir -p "/home/bin"

COPY ./bin/server /home/bin/

RUN chmod -R 777 /home

WORKDIR /home/bin

ENTRYPOINT ["./server"]
CMD ["server", "-c", "server.toml"]
