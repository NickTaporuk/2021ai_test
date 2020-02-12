package stack

import (
	"errors"
)

var (
	// ErrorStackIsEmpty use for notify about the structure stack parameter slice is empty
	ErrorStackIsEmpty = errors.New("can't pop; stack is empty")
	// ErrorNoElements use for notify about the stack doesn't have element inside
	ErrorNoElements = errors.New("no elements in stack")
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
		return "", ErrorStackIsEmpty
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
		return "", ErrorNoElements
	}
	return s.Slice[s.Pos], nil
}

func (s *StringStack) SafeTop() string {
	ret, _ := s.Top()
	return ret
}
