package main

import (
	"test.com/test"
)

_#TestName: =~"test"

test1: test.#Test & {
	name:   _#TestName & "test 1"
	assert: "it's the first test"
}

test2: test.#Test & {
	name:   _#TestName & "test 2"
	assert: "it's the second test"
}
