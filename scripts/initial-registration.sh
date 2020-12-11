#!/usr/bin/env bash
java -jar /opt/corda/bin/corda.jar \
        initial-registration \
        --config-file $1 \
        --network-root-truststore-password=$2 \
        --network-root-truststore=/opt/corda/truststore.jks \
        --base-directory $3