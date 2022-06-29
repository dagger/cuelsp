package semantic

import (
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

// sums every 4th element of the array in reverse while sum - cost >= 0
// returns the splitIndex => Index at which the array should be split
// returns the leftOver => distance from previous token to the one we will append
// returns the distance => distance from appended token to ones we will move
func indexFinder(arr []protocol.UInteger, cost int) (int, int, int, bool) {
	sum := 0
	for i := len(arr) - 4; i >= 0; i -= 5 {
		sum += int(arr[i])
		if cost-sum <= 0 {
			splitIndex := i - 1
			leftOver := sum - cost
			distance := int(arr[i]) - leftOver
			return splitIndex, leftOver, distance, true
		}
	}
	return -1, -1, -1, false
}

// appends an array in arr at index
func appendArrayAtIndex(arr []protocol.UInteger, secondArray []protocol.UInteger, index uint32) []protocol.UInteger {
	var newArr []protocol.UInteger
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, secondArray...)
	newArr = append(newArr, arr[index:]...)
	return newArr
}

// // Update distance for moved token to appended token
func updateDistanceAtIndex(arr []protocol.UInteger, index uint32, value uint32) []protocol.UInteger {
	arr[index] = value
	return arr
}

// untested
func shiftTokens(tokens []protocol.UInteger, newToken []protocol.UInteger, cost int, logger logging.Logger) ([]protocol.UInteger, bool) {
	index, leftOver, distance, ok := indexFinder(tokens, cost)
	if !ok || index < 0 || leftOver < 0 || distance < 0 {
		logger.Debugf("crap %d %d %d %d", ok, index, leftOver, distance)
		logger.Debugf("crap2 %v %d", tokens, cost)
		return nil, false
	}
	logger.Debugf("it passed")
	tokens = updateDistanceAtIndex(tokens, uint32(index+1), uint32(distance))
	newToken = updateDistanceAtIndex(newToken, 1, uint32(leftOver))
	newArr := appendArrayAtIndex(tokens, newToken, uint32(index))
	return newArr, true
}

// untested
func newToken(patternIndex uint16, n *sitter.Node) []protocol.UInteger {
	return []uint32{0, 0, n.EndPoint().Column - n.StartPoint().Column, tokenTypeIndex(patternIndex), tokenModifierIndex(patternIndex)}
}
