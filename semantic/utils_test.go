package semantic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Find Index to update according to cost
// asserts index, leftOver and distance
func TestIndexFinder(t *testing.T) {
	arr := []uint32{
		1, 4, 0, 0, 0,
		1, 5, 0, 0, 0,
		1, 2, 0, 0, 0,
	}
	cost := 5
	index, leftOver, distance, ok := indexFinder(arr, cost)
	if !ok {
		require.Failf(t, "IndexFinder failed", "index: %d, leftOver: %d, distance: %d", index, leftOver, distance)
	}
	require.Equal(t, 5, index)    // index to send to AppendArrayAtIndex
	require.Equal(t, 2, leftOver) // leftOver for secondArray in AppendArrayAtIndex
	require.Equal(t, 3, distance) // distance is value in UpdateDistanceAtIndex
}

// Append array in array at index
func TestAppendArrayAtIndex(t *testing.T) {
	arr := []uint32{1, 2, 3, 4, 5, 6, 7, 8}
	secondArray := []uint32{1, 2, 3, 4, 5, 6, 7, 8}

	newArr := appendArrayAtIndex(arr, secondArray, uint32(4))
	require.Equal(t, []uint32{1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 5, 6, 7, 8}, newArr)
}

// func TestUpdateDistanceAtIndex(t *testing.T) {
// 	arr := []uint32{1, 2, 3, 4}

// 	newArr := updateDistanceAtIndex(arr, 1, uint32(5))
// 	require.Equal(t, []uint32{1, 5, 3, 4}, newArr)
// }
