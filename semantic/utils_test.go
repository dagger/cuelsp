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
	index, leftOver, distance, ok := indexFinder(arr, float64(cost))
	if !ok {
		require.Failf(t, "IndexFinder failed", "index: %d, leftOver: %d, distance: %d", index, leftOver, distance)
	}
	require.Equal(t, 5, index)    // index to send to AppendSliceAtIndex
	require.Equal(t, 2, leftOver) // leftOver for secondSlice in AppendSliceAtIndex
	require.Equal(t, 3, distance) // distance is value in UpdateDistanceAtIndex
}

// Appends slice to slice, at index
func TestAppendSliceAtIndex(t *testing.T) {
	arr := []uint32{1, 2, 3, 4, 5, 6, 7, 8}
	secondSlice := []uint32{1, 2, 3, 4, 5, 6, 7, 8}

	newArr := appendSliceAtIndex(arr, secondSlice, 4)
	require.Equal(t, []uint32{1, 2, 3, 4, 1, 2, 3, 4, 5, 6, 7, 8, 5, 6, 7, 8}, newArr)
}

// Updates distance at index 1 to value 5
func TestUpdateDistanceAtIndex(t *testing.T) {
	arr := []uint32{1, 2, 3, 4}

	newArr := updateDistanceAtIndex(arr, 1, 5)
	require.Equal(t, []uint32{1, 5, 3, 4}, newArr)
}

func TestShiftTokens(t *testing.T) {
	tokens := []uint32{
		0, 2, 4, 1, 0, // first token
		0, 7, 4, 2, 0, // second token
	}
	newToken := []uint32{0, 4, 4, 3, 0} // token to append in between

	newArr, ok := shiftTokens(tokens, newToken, 5, nil)
	require.True(t, ok)
	require.Equal(t, []uint32{
		0, 2, 4, 1, 0, // first token not shifted
		0, 2, 4, 3, 0, // new token appended here
		0, 5, 4, 2, 0, // previous token updated
	}, newArr)
}
