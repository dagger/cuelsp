package ci

import (
	"dagger.io/dagger"

	"universe.dagger.io/go"
	"universe.dagger.io/alpha/go/golangci"
)

dagger.#Plan & {
	// Input
	client: filesystem: ".": read: {
		include: ["**/*.go", "go.mod", "go.sum", ".golangci.yaml", "**/testdata/*", ".golangci.yaml"]
	}

	// Output
	client: filesystem: {
		"/tmp/cov.html": write: {
			contents: actions.test.coverage.export.files."/tmp/cov.html"
		}
		"/tmp/cov.txt": write: {
			contents: actions.test.coverage.export.files."/tmp/cov.txt"
		}
		"./bin": write: {
			contents: actions.build.output
		}
	}

	actions: {
		_code: client.filesystem.".".read.contents

		build: go.#Build & {
			source:  _code
			package: "github.com/dagger/cuelsp/cmd/cuelsp"
		}

		test: {
			go.#Test & {
				source:  _code
				package: "./..."
				command: flags: {
					"-race":         true
					"-coverprofile": "/tmp/cov.txt"
				}
			}

			coverage: go.#Container & {
				image:  test.output
				source: _code
				command: {
					name: "sh"
					args: ["-c", """
						go tool cover -html=/tmp/cov.txt -o /tmp/cov.html
						"""]
				}
				export: files: {
					"/tmp/cov.html": string
					"/tmp/cov.txt":  string
				}
			}
		}

		lint: golangci.#Lint & {
			source: _code
		}
	}
}
