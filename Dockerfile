FROM alpine:3.19.1 as build
RUN apk add --no-cache go && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY ./app .
RUN go mod init app && \
    go build -o main ./main.go

FROM alpine:3.19.1 as prod
USER 1000:1000
COPY --from=build /app/main .
CMD ["./main"]