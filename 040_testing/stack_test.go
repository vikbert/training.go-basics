package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceStack_IsEmpty(t *testing.T) {
	s := SliceStack{}
	assert.Empty(t, s, "Init stack should be empty")
}

func TestSliceStack_Push(t *testing.T) {
	s := SliceStack{}
	s.Push(12)
	size := s.Size()
	assert.Equal(t, 1, size, "Stack should have the length '1'")
}

func TestSliceStack_Peek(t *testing.T) {
	s := SliceStack{}
	s.Push(12)
	got := s.Peek()
	assert.Equal(t, 12, got, "Stack should return the expected value")
}

func TestSliceStack_Pop(t *testing.T) {
	s := SliceStack{}
	s.Push(12)
	got := s.Pop()
	assert.Equal(t, 12, got, "Stack should return the expected value")
	assert.Empty(t, s, "Stack should be empty after remove the single element")
}
