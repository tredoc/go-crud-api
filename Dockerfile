FROM golang:1.22-alpine

WORKDIR /app
COPY ./bin/main .
COPY ./.env .
CMD ["/app/main"]