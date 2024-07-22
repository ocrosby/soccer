#!/usr/bin/env zsh

pushd ../../
# Generate the Swagger API documentation for the tds-coaching-change-service
swagger generate spec -o ./services/tds-coaching-change-service/swagger.yaml --scan-models ./internal/tds-coaching-change-service ./services/tds-coaching-change-service ./pkg
popd
