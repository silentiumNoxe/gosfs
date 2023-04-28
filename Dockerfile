FROM golang:1.18

WORKDIR /app

COPY ./gosfs gosfs
COPY ./public public
COPY ./gosfs.json gosfs.json

EXPOSE 8080

ENTRYPOINT ["/app/gosfs"]