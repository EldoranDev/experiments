package stack

import "errors"

type node struct {
	val  uint16
	prev *node
}

type Stack struct {
	top    *node
	length int
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() (uint16, error) {
	if s.length == 0 {
		return 0, errors.New("there is nothing on the Stack")
	}

	return s.top.val, nil
}

func (s *Stack) Pop() (uint16, error) {
	if s.length == 0 {
		return 0, errors.New("there is nothing on the Stack")
	}

	node := s.top
	s.top = node.prev

	s.length--

	return node.val, nil
}

func (s *Stack) Push(value uint16) {
	node := &node{value, s.top}

	s.top = node
	s.length++
}
