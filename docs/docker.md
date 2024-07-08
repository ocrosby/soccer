# Docker

This document contains notes on Docker.

Building a Docker image for a specific package

```shell
make build SERVICE=tds-college-conference ENV=dev
```

Notice that the "-service" part of the name is omitted.

Since the Docker build context is at the root of the repository ...

Checking the size of the Docker image

```shell
docker images tds-college-conference-service:latest
```

## Keeping Docker Clean

```shell
docker image prune -f
docker container prune -f
docker builder prune -f
docker system prune -a -f
```

List all dangling images

```shell
docker images -f dangling=true -q
```

Delete all dangling images

```shell
docker images -f dangling=true -q | xargs docker rmi
```