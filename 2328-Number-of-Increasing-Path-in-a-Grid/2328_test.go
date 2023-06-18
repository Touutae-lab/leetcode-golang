package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePathsIII(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(8, countPaths([][]int{{1, 1}, {3, 4}}))
	assert.Equal(3, countPaths([][]int{{1}, {2}}))
}
