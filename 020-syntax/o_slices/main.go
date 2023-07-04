package main

import "fmt"

type SliceStack struct {
	Data []interface{}
}

func main() {
	myStack := SliceStack{}
	myStack.Push(0)
	myStack.Push(1)
	myStack.Push(2)

	fmt.Println("Slices:", myStack)
	fmt.Println("Slices::Pop:", myStack.Pop())
	fmt.Println("Slices:", myStack)
	fmt.Println("Slices::Peek:", myStack.Peek())
	fmt.Println("Slices:", myStack)
	fmt.Println("Slices::Size:", myStack.Size())
	fmt.Println("Slices::IsEmpty:", myStack.IsEmpty())
}

func (s *SliceStack) Push(value interface{}) {
	s.Data = append(s.Data, value)
}

func (s *SliceStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	value := s.Data[0]
	s.Data = s.Data[1:len(s.Data)]

	return value
}

func (s *SliceStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.Data[0]
}

func (s *SliceStack) Size() int {
	return len(s.Data)
}

func (s *SliceStack) IsEmpty() bool {
	return len(s.Data) == 0
}
