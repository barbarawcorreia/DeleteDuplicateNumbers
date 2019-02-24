package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesDuplicatedSequency(t *testing.T) {
	sliceDuplicated := []int{2, 2, 3, 3, 4, 4, 5, 5}
	sliceExpeted := []int{2, 3, 4, 5}

	sliceReturned := DeleteDuplicated(sliceDuplicated)
	assert.Equal(t, sliceExpeted, sliceReturned)
}

func TestSlicesDuplicatedNotInSequency(t *testing.T) {
	sliceDuplicated := []int{3, 2, 2, 3, 6}
	sliceExpeted := []int{2, 3, 6}

	sliceReturned := DeleteDuplicated(sliceDuplicated)
	assert.Equal(t, sliceExpeted, sliceReturned)
}
