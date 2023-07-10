package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setup() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	teardown()
	os.Exit(exitVal)
}

func TestSliceStack_IsEmpty(t *testing.T) {
	s := SliceStack{}
	assert.True(t, s.IsEmpty(), "Init stack should be empty")
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
	assert.True(t, s.IsEmpty(), "Stack should be empty after remove the single element")
}
