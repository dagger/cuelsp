package dir

import (
	"test.com/test"
)

tmp: #Path & "./test"

test1: test.#Test & {
	name:   tmp
	assert: "it's the first test"
}
