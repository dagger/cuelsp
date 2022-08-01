package main

import (
	"test.com/test"
	t "test.com/test2"
)

_#TestName: =~"test"

test1: test.#Test & {
	name:   _#TestName & "test 1"
	assert: "it's the first test"
}

test2: t.#Test & {
	name:   _#TestName & "test 2"
	assert: "it's the second test"
}
