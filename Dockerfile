FROM golang:1.15.0 AS builder

COPY . /src/go/
WORKDIR /src/go/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd/server/main.go


FROM adoptopenjdk/openjdk8

RUN curl -LJO https://dl.bintray.com/r3/corda/net/corda/corda/4.6/corda-4.6.jar \
    && mkdir -p /opt/corda/bin \
    && mv corda-4.6.jar /opt/corda/bin/corda.jar

COPY ./scripts .
RUN mv initial-registration.sh /bin/initial-registration \
    && mv validate-configuration.sh /bin/validate-configuration \
    && mv truststore.jks /opt/corda/truststore.jks

COPY ./config /config
COPY --from=builder /src/go/app .

EXPOSE 8080

ENV STOAGE_PATH /data/corda
ENV CONFIG_PATH /config

ENTRYPOINT ["./app"]