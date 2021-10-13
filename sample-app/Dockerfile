FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o /bin/app

FROM alpine:3.14

COPY --from=builder /bin/app /bin/app

EXPOSE 8080
EXPOSE 20777/udp

ENTRYPOINT ["/bin/app"]
