FROM golang:alpine as builder
RUN apk update && apk add git
RUN adduser -D -g '' cloud2podcast

COPY . $GOPATH/src/c2p
WORKDIR $GOPATH/src/c2p

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s'  -o /go/bin/cloud2podcast

# Build small sratch image
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/cloud2podcast /go/bin/cloud2podcast
USER cloud2podcast
VOLUME /downloads
EXPOSE 8080
ENTRYPOINT ["/go/bin/cloud2podcast"]