package pipefilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sum)

	ret, err := sp.Process("1,2,3")
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, ret)
}
