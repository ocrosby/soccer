#!/usr/bin/env zsh

pushd ../../
# Run the Docker image for the tds-coaching-change-service
docker run -d --name tds-coaching-change-service -p 8080:8080 tds-coaching-change-service
popd
