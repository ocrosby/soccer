# Soccer

This repository contains a monorepo for Soccer related projects.

## Usage

When someone clones your monorepo, they will need to rebuild the go.work file as follows:

```shell
go work init
go work use -r .
```

### Creating a package

To create a new package

```shell
mkdir -p ./pkg/<package-name>
```

Initialize the package with a go.mod file

```shell
go mod init -C ./pkg/<package-name> github.com/ocrosby/soccer/pkg/<package-name>
```

Add the new package to the go.work file

```shell
go work use ./pkg/<package-name>
```

Creating a tag

```shell
git tag v0.1.0
git push origin --tags
```

Every time you publish new changes, run that again, updating the version number (0.1.0, 0.1.1, 0.1.2, etc.)

## Project Structure

In a monorepo setup for Go projects, whether to use a single go.mod at the root or individual go.mod files in each cmd service directory depends on the structure and dependencies of your projects. Here are considerations for both approaches:  
Single go.mod at the Root

Simplicity: A single go.mod file simplifies dependency management by having a unified set of dependencies for all projects within the monorepo.

Consistency: Ensures all projects within the monorepo use the same versions of dependencies, which can help in avoiding compatibility issues.

Tooling: Simplifies tooling and scripting around builds, tests, and dependency updates.
Individual go.mod Files in Each cmd Service Directory

Isolation: Allows each service to manage its dependencies independently, which is useful if the services are intended to be deployed or versioned separately.

Flexibility: Services can upgrade their dependencies independently, reducing the risk of conflicts or forced updates across unrelated services.

Microservice Architecture: Fits well with a microservice architecture where services are loosely coupled and may have different lifecycles and dependency requirements.

### Decision Factors

Project Size and Complexity: For smaller or tightly coupled projects, a single go.mod may be sufficient. For larger, more complex projects with loosely coupled services, individual go.mod files may offer better flexibility and isolation.

Dependency Management: Consider whether you prefer centralized management of dependencies or if you need the flexibility to manage dependencies per service.

Deployment and Versioning: If services are deployed and versioned independently, separate go.mod files might be more appropriate.

it Given the structure of your monorepo and the presence of a cmd directory that suggests potentially multiple services, if these services are intended to be independently deployable or have significantly different dependencies, it might be beneficial to have individual go.mod files for each service. This approach provides greater flexibility and isolation for dependency management, which can be advantageous in a microservices architecture.

## BDD Testing with Godog

Install the latest version of the Godog CLI

```shell
go install github.com/cucumber/godog/cmd/godog@latest
```

Run the tests

```shell
godog
```

## Docker Build Issues

Given your monorepo structure and the details provided, the best approach for building Docker images involves a few key steps to ensure efficiency, maintainability, and scalability. Here's a detailed plan:  

1. Centralize Docker Build Context: Use the root of the monorepo as the Docker build context. This allows all Dockerfiles, regardless of their location within the repository, to access the entire codebase and any shared resources or configurations.  

2. Parameterize Build Arguments: Utilize build arguments (--build-arg) in your Docker build commands to dynamically specify environment configurations, versions, or any other necessary parameters. This enhances flexibility across different environments (development, testing, production) without hardcoding values.  

3. Optimize Dockerfile for Caching: Structure your Dockerfile to leverage Docker's layer caching mechanism effectively. Copy and install dependencies before copying the entire application code to avoid unnecessary rebuilds of dependencies.  

4. Multi-Stage Builds: Implement multi-stage builds in your Dockerfile to minimize the final image size by separating the build environment from the runtime environment. This approach allows you to include only the compiled binaries and necessary runtime dependencies in the final image.  

5. Consistent Directory Structure: Maintain a consistent directory structure within your Docker images. This simplifies the process of running the containers and managing volumes or configurations, especially when dealing with multiple services or applications within the monorepo.  

6. Documentation and Comments: Document the build process and any specific commands or configurations directly within the Dockerfile and accompanying documentation. This is crucial for onboarding new developers and ensuring clarity across the team.  

7. Automate Builds with CI/CD: Integrate Docker image building and pushing into your CI/CD pipeline. Use scripts or CI/CD pipeline configurations to automate the building of images based on changes to specific directories or files within the monorepo.  

8. Versioning and Tagging Strategy: Implement a consistent versioning and tagging strategy for your Docker images. This could involve using semantic versioning, git commit hashes, or environment-specific tags to easily identify and rollback to specific image versions if necessary.





Building the tds-college-conference-service image

```shell
make build SERVICE=tds-college-conference ENV=dev
```

Running the tds-college-conference-service container

```shell
make run SERVICE=tds-college-conference ENV=dev
```

## Running Tests

To run the integration tests for location-service

```shell
go test -tags=godog ./cmd/location-service/features
```

Building the location service image

```shell
docker build -f cmd/location-service/Dockerfile -t location-service:latest .
```

## Dealing with Docker

Cleaning up dangling images

```shell
docker image prune
```

## References

- [Test Containers for Go](https://golang.testcontainers.org/)
- [Project Documentation](docs/index.md)
- [How to use Godog for BDD development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go)
- [Example Go Monorepo](https://github.com/Medium-Stories/go-mono-repo)
- [Go Project Structure Monorepo](https://blog.devops.dev/go-project-structure-monorepo-daa762ec36a2)
- [How to create and use a Go monorepo](https://scriptable.com/golang/how-to-create-and-use-a-go-monorepo/)
- [Getting started with multi-module workspaces in Go](https://go.dev/doc/tutorial/workspaces/)
- [Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)
- [How I write HTTP services in Go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
- [Local Development of Go Applications with Testcontainers](https://www.docker.com/blog/local-development-of-go-applications-with-testcontainers/)
- [Logging HTTP Requests in Go](https://blog.kowalczyk.info/article/e00e89c3841e4f8c8c769a78b8a90b47/logging-http-requests-in-go.html)