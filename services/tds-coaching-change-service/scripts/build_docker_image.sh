#!/usr/bin/env zsh

pushd ../../
docker build -t tds-coaching-change-service -f services/tds-coaching-change-service/Dockerfile .
docker image prune -f
popd