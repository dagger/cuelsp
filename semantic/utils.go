package semantic

import (
	"io/ioutil"
	"math"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

// sums every 4th element of the slice in reverse while sum - cost >= 0
// returns the splitIndex => Index at which the slice should be split
// returns the leftOver => distance from previous token to the one we will append
// returns the distance => distance from appended token to ones we will move
func indexFinder(arr []protocol.UInteger, cost float64) (int, int, int, bool) {
	sum := 0
	absCost := int(math.Abs(cost))

	for i := len(arr) - 4; i >= 0; i -= 5 {
		sum += int(arr[i])
		if absCost-sum <= 0 {
			splitIndex := i - 1
			leftOver := sum - absCost
			distance := int(arr[i]) - leftOver
			return splitIndex, leftOver, distance, true
		}
	}
	return -1, -1, -1, false
}

// appends an slice in arr at index
func appendSliceAtIndex(arr []protocol.UInteger, secondSlice []protocol.UInteger, index int) []protocol.UInteger {
	var newArr []protocol.UInteger
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, secondSlice...)
	newArr = append(newArr, arr[index:]...)
	return newArr
}

// Update distance for moved token to appended token
func updateDistanceAtIndex(arr []protocol.UInteger, index int, value int) []protocol.UInteger {
	arr[uint32(index)] = uint32(value)
	return arr
}

// Tested manually
// Find index to insert token to append, and where to start shifting the others on the right
func shiftTokens(tokens []protocol.UInteger, newToken []protocol.UInteger, cost float64, logger logging.Logger) ([]protocol.UInteger, bool) {
	index, leftOver, distance, ok := indexFinder(tokens, cost)
	if !ok || index < 0 || leftOver < 0 || distance < 0 {
		logger.Errorf("shiftTokens Error: %d %d %d %d %v %d", ok, index, leftOver, distance, tokens, cost)
		return nil, false
	}

	// Update token.Column() to future pos after shift of tokens
	tokens = updateDistanceAtIndex(tokens, index+1, distance)
	// Update newtoken.Column position to distance from previous token
	newToken = updateDistanceAtIndex(newToken, 1, leftOver)
	// Append newToken to token slice
	newArr := appendSliceAtIndex(tokens, newToken, index)
	return newArr, true
}

func readFile(filename string, logger logging.Logger) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Errorf("readFile Error: %s", err)
		return nil, err
	}
	return data, nil
}

// Checks whether a token is a is multiline string
func isMultilineStringToken(node string) bool {
	return strings.HasPrefix(node, "(multiline_string_lit")
}
