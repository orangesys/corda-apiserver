#!/usr/bin/env bash
java -jar $1/corda.jar \
        --initial-registration \
        --config-file $2/node.conf \
        --network-root-truststore-password=trustpass \
        --network-root-truststore=./truststore.jks \
        --base-directory /Users/xucheng/go/src/corda-apiserver/tmp/efb0c9745bc1b98b05bd6c9e93e91c35
echo "Succesfully registered"