# CUE LSP

Language Server implementation for [CUE](https://cuelang.org), with built-in support for [Dagger](https://dagger.io).

## Install

CUE LSP can be used in one of two ways:

* As a standalone binary `cuelsp`: `go install github.com/dagger/cuelsp/cmd/cuelsp@latest`
* As a subcommand of the Dagger CLI: `dagger cuelsp`. This requires [dagger v0.2.27 or later](https://docs.dagger.io/install).

## Configure your code editor

| Code Editor |                    Documentation                     |
|:-----------:|:----------------------------------------------------:|
|   VS Code   | [Extension](https://marketplace.visualstudio.com/items?itemName=Dagger.dagger) |
|     Vim     |                [Guide](./docs/vim.md)                |

### Capabilities

| Feature                 | Supported          | Link to documentation                     |
|-------------------------|--------------------|-------------------------------------------|
| Load cue plan           | :white_check_mark: | [how cuelsp loads CUE](./docs/load.md) |
| Load multiples files    | :white_check_mark: | [how cuelsp loads CUE](./docs/load.md) |
| Jump to CUE definition  | :white_check_mark: | [managing jump-to](./docs/jump-to.md)       |
| Doc Hover               | :white_check_mark: |                                           |
| Syntax highlighting     | :hourglass:        |                                           |
| Auto completion         | :no_entry_sign:    |                                           |
| Jump to CUE keys        | :no_entry_sign:    |                                           |
| Error highlighting      | :no_entry_sign:    |                                           |
| Code snippet            | :no_entry_sign:    |                                           |
| Optimization suggestion | :no_entry_sign:    |                                           |

### Development & CI

We use [Dagger](https://dagger.io) to lint, test and build CUELSP. Using Dagger, commands running in the
CI behave the same as on your local system :rocket:

| Action       | Command           |
|--------------|-------------------|
| Run linter   | `dagger do lint`  |
| Run test     | `dagger do test`  |
| Build binary | `dagger do build` |

> If you are on Mac M1, you should build binary using `go build -o daggerlsp` because Buildkit
> does not support `darwin/arm64` platform, yet.

### Release

CUELSP is versioned through tagged release.

There is a complete [release workflow](./.github/workflows/release.yaml) to produce CUELSP binaries for multiple
platforms.

To publish a new release, we just create a new tag.

```shell
# Tag current commit
git tag vX.X.X

# Push tag to repository
git push origin vX.X.X
```

### Maintainers

| [<img src="https://github.com/TomChv.png?size=85" /><br /><sub><b>Vasek</b></sub>](https://github.com/TomChv) | [<img src="https://github.com/grouville.png?size=85" /><br /><sub><b>Guillaume de Rouville</b></sub>](https://github.com/grouville) | [<img src="https://github.com/dolanor.png?size=85" /><br /><sub><b>Tanguy ⧓ Herrmann</b></sub>](https://github.com/dolanor) |
|:-------------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------------------------------:|:---------------------------------------------------------------------------------------------------------------------------:|
