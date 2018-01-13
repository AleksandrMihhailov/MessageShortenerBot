FROM golang

WORKDIR /go/src/MessageShortenerBot
COPY . .
RUN go get && go build

ENTRYPOINT [ "/go/src/MessageShortenerBot/MessageShortenerBot" ]
