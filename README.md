# Dagger LSP

Language Server implementation for [Dagger](https://github.com/dagger/dagger).

## Install

### Releases

Get the [latest released version](https://github.com/grouville/dagger-lsp/releases/latest) and put it in your `$PATH`

#### Linux
```
curl -L https://github.com/grouville/dagger-lsp/releases/download/v0.1.0/daggerlsp_0.1.0_linux_amd64.tar.gz | tar -xzC .
sudo install daggerlsp /usr/local/bin/
```

### Latest from source

You need to have [Go installed](https://go.dev/doc/install), and your `$(go env GOPATH)/bin` path should be in your `$PATH`.

```
go install github.com/grouville/dagger-lsp@latest
```

## Use in your IDE

### vim

[Documentation](/docs/vim.md)

### Development & CI

Current CI is using [Dagger](https://dagger.io) to lint, test and build the LSP.  
Using Dagger, commands running in the CI behave the same as on your local system :rocket:

| Action       | Command           |
|--------------|-------------------|
| Run linter   | `dagger do lint`  |
| Run test     | `dagger do test`  |
| Build binary | `dagger do build` |

> If you are on Mac M1, you should build binary using `go build -o dlsp` because Buildkit
> does not support `darwin/arm64` platform.

### Capabilities

| Feature                 | Supported          | Link to documentation               |
|-------------------------|--------------------|-------------------------------------|
| Load cue plan           | :white_check_mark: | [how dlsp load CUE](./docs/load.md) |
| Load multiples files    | :white_check_mark: | [how dlsp load CUE](./docs/load.md) |
| Jump to CUE definition  | :white_check_mark: | [manage jump-to](./docs/jump-to.md) |
| Syntax highlighting     | :hourglass:        |                                     |
| Doc Hover               | :no_entry_sign:    |                                     |
| Auto completion         | :no_entry_sign:    |                                     |
| Jump to CUE keys        | :no_entry_sign:    |                                     |
| Error highlighting      | :no_entry_sign:    |                                     |
| Code snippet            | :no_entry_sign:    |                                     |
| Optimization suggestion | :no_entry_sign:    |                                     |

### Release

Dagger LSP is versioned through tagged release.

There is a complete [release workflow](./.github/workflows/release.yaml) to populate Dagger LSP binary in multiple
platforms.

To publish a new release, just create a new tag

```shell
# Tag current commit
git tag vX.X.X

# Push tag to repository
git push origin vX.X.X
```

### Maintainers

| [<img src="https://github.com/TomChv.png?size=85" /><br /><sub><b>Vasek</b></sub>](https://github.com/TomChv) | [<img src="https://github.com/grouville.png?size=85" /><br /><sub><b>Guillaume de Rouville</b></sub>](https://github.com/grouville) | [<img src="https://github.com/dolanor.png?size=85" /><br /><sub><b>Tanguy ⧓ Herrmann</b></sub>](https://github.com/dolanor) |
| :---: | :---: | :---: |
