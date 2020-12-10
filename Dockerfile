FROM golang:1.15.0 AS builder

COPY . /src/go/
WORKDIR /src/go/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd/server/main.go


FROM alpine
WORKDIR /root/
COPY --from=builder /src/go/app .

EXPOSE 8080
ENTRYPOINT ["./app"]