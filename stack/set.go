package stack

import "errors"

// SetStacks
type SetStack struct {
	Slice [][]int
	Pos   int
}

//
func NewSetStack() SetStack {
	return SetStack{
		Slice: [][]int{},
		Pos:   -1,
	}
}

// Push
func (s *SetStack) Push(a []int) {
	s.Pos++
	if s.Pos < len(s.Slice) {
		s.Slice[s.Pos] = a
	} else {
		s.Slice = append(s.Slice, a)
	}
}

// Pop
func (s *SetStack) Pop() ([]int, error) {
	ret, err := s.Top()

	if err != nil {
		return []int{}, errors.New("Can't pop; stack is empty!")
	}

	s.Pos--

	return ret, nil
}

func (s *SetStack) Top() ([]int, error) {
	if s.Pos < 0 {
		return []int{}, errors.New("No elements in stack!")
	}
	return s.Slice[s.Pos], nil
}
