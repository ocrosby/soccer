#!/usr/bin/env zsh

# Stop the Docker container for the tds-coaching-change-service
docker stop tds-coaching-change-service

# Optionally, remove the container after stopping
docker rm tds-coaching-change-service