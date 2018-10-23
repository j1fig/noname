FROM golang:1.11-alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
RUN mkdir $HOME/.ssh
RUN echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > $HOME/.ssh/config
RUN mkdir /build 
ADD . /build/
RUN mv /build/keys/build_id_rsa $HOME/.ssh/id_rsa
RUN mv /build/keys/build_id_rsa.pub $HOME/.ssh/id_rsa.pub
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
