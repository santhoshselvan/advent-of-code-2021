package adventOfCode2021

import "testing"
import "github.com/stretchr/testify/assert"

func Test_findNumberOfIncreasingSequence(t *testing.T) {
	assert.Equal(t, 0, FindNumberOfIncreasingSequence([]string{"199"}))
	assert.Equal(t, 1, FindNumberOfIncreasingSequence([]string{"1", "2"}))
	assert.Equal(t, 1, FindNumberOfIncreasingSequence([]string{"1", "2", "1"}))
	assert.Equal(t, 2, FindNumberOfIncreasingSequence([]string{"1", "2", "1", "5"}))
	assert.Equal(t, 0, FindNumberOfIncreasingSequence([]string{"7", "5", "3", "1"}))
	assert.Equal(t, 7,
		FindNumberOfIncreasingSequence([]string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}))
}
