FROM golang:1.13-alpine as builder
# ENV GO111MODULE=on
WORKDIR /go/src/app
RUN apk update && apk add --no-cache git
COPY ./cmd/main.go .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo -o /go/bin/app .

FROM scratch
COPY --from=builder /go/bin/app /app
EXPOSE 9000
ENTRYPOINT ["/app"]