FROM golang:1.17 as builder

ENV GOPATH="/goroot/"
WORKDIR /goroot/src/github.com/sirnarsh/gotelltherabbit/
COPY . /goroot/src/github.com/sirnarsh/gotelltherabbit/
RUN ls
RUN go get
RUN CGO_ENABLED=0 go build

FROM alpine:latest
# mailcap adds mime detection and ca-certificates help with TLS (basic stuff)
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /goroot/src/github.com/sirnarsh/gotelltherabbit/gotelltherabbit /app/gotelltherabbit
VOLUME /app/config/

EXPOSE 80
EXPOSE 8080
ENTRYPOINT ["./gotelltherabbit"]
