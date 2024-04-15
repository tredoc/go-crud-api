FROM alpine:3.14

WORKDIR /app
COPY ./bin/main .
COPY ./.env .
CMD ["/app/main"]