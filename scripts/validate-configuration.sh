#!/usr/bin/env bash
java -jar /opt/corda/bin/corda.jar \
        validate-configuration \
        --config-file $1 \
        --base-directory $2