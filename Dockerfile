FROM golang:1.11-alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go get -d -v ./...
RUN go build -o main .

FROM alpine
RUN apk add --no-cache ca-certificates apache2-utils
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
COPY --from=builder /build/data/ /app/data/
WORKDIR /app
CMD ["./main"]
