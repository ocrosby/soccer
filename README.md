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


## References

- [How to create and use a Go monorepo](https://scriptable.com/golang/how-to-create-and-use-a-go-monorepo/)
- [Getting started with multi-module workspaces in Go](https://go.dev/doc/tutorial/workspaces/)
- [Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)