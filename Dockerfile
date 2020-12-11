FROM golang:1.15.0 AS builder

COPY . /src/go/
WORKDIR /src/go/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd/server/main.go


FROM adoptopenjdk/openjdk8

COPY ./corda.jar /opt/corda/bin/corda.jar
COPY ./initial-registration.sh /bin/initial-registration
COPY ./validate-configuration.sh /bin/validate-configuration
COPY ./truststore.jks /opt/corda/truststore.jks
COPY ./config /config
COPY --from=builder /src/go/app .

EXPOSE 8080

ENV STOAGE_PATH /data/corda
ENV CONFIG_PATH /config

ENTRYPOINT ["./app"]