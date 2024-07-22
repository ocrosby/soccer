# Swagger

I will be utilizing [go-swagger](https://github.com/go-swagger/go-swagger) in this project.


## Installation

```bash
go install github.com/go-swagger/go-swagger/services/swagger@latest
```

Note: Make sure you have `$GOPATH/bin` in your `$PATH`.

## Generate Swagger Docs

```bash
swagger generate spec -o ./docs/swagger.json
```

Note: `swagger generate specs` operates on the current directory and it's subdirectories.

You will need to ensure you are in the root directory of your specific service within the monorepo.

Use the -m flag to tell swagger to only include Go files that are part of the main package.


Generalized command structure:

```bash
swagger generate spec -o ./docs/swagger.json -m [directories] --exclude [excluded_directories]
```

