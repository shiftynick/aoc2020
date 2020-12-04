package day4

import (
	"testing"

	"github.com/shiftynick/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := a(input)

	require.Equal(t, 42, result)
}
