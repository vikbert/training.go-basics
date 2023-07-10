package stack

type SliceStack struct {
	Data []interface{}
}

func NewStack() *SliceStack {
	return &SliceStack{}
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
