# Coaching Change Service Tests

To run the integration tests for location-service, from the root directory execute the following

```shell
go test -tags=godog ./services/tds-coaching-change-service/features
```

Build the Docker Image for location-service

```shell
docker build -t tds-coaching-change-service -f services/tds-coaching-change-service/Dockerfile .
```