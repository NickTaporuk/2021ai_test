package stack

import (
	"errors"
)

type StringStack struct {
	Slice []string
	Pos   int
}

func NewStringStack() *StringStack {
	return &StringStack{
		Slice: []string{},
		Pos:   -1,
	}
}

func (s *StringStack) Push(a string) {
	s.Pos++
	if s.Pos < len(s.Slice) {
		s.Slice[s.Pos] = a
	} else {
		s.Slice = append(s.Slice, a)
	}
}

func (s *StringStack) Pop() (string, error) {
	ret, err := s.Top()
	if err != nil {
		return "", errors.New("Can't pop; stack is empty!")
	}
	s.Pos--
	return ret, nil
}

func (s *StringStack) SafePop() string {
	ret, _ := s.Pop()
	return ret
}

func (s *StringStack) Top() (string, error) {
	if s.Pos < 0 {
		return "", errors.New("No elements in stack!")
	}
	return s.Slice[s.Pos], nil
}

func (s *StringStack) SafeTop() string {
	ret, _ := s.Top()
	return ret
}
