package adventOfCode2021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberOfIncreasingSlidingWindow(t *testing.T) {

	assert.Equal(t, 0, FindNumberOfIncreasingSlidingWindow([]string{"199", "200", "208"}))
	assert.Equal(t, 0, FindNumberOfIncreasingSlidingWindow([]string{"1", "2", "3", "1"}))
	assert.Equal(t, 1, FindNumberOfIncreasingSlidingWindow([]string{"1", "2", "3", "2"}))
	assert.Equal(t, 1, FindNumberOfIncreasingSlidingWindow([]string{"1", "2", "3", "2", "2"}))
	assert.Equal(t, 2, FindNumberOfIncreasingSlidingWindow([]string{"1", "2", "3", "2", "4"}))
	assert.Equal(t, 5,
		FindNumberOfIncreasingSlidingWindow([]string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}))
}
