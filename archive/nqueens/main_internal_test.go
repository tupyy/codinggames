package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckQSolution(t *testing.T) {
	b := [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}

	assert.False(t, checkQSolution(b, 2, 1))
}
func TestCheckQSolution2(t *testing.T) {
	b := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, checkQSolution(b, 2, 1))
}

func TestCheckQSolution3(t *testing.T) {
	b := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	}

	assert.False(t, checkQSolution(b, 0, 0))
}

func TestCheckQSolution4(t *testing.T) {
	b := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.True(t, checkQSolution(b, 0, 0))
}

func TestCheckQSolution5(t *testing.T) {
	b := [][]int{
		{1, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	assert.False(t, checkQSolution(b, 0, 3))
}
