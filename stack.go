package main

import (
	"errors"
)

type variableInfo struct {
	Name  string
	Value int
}

type codeBlockVariableInfo struct {
	Info  variableInfo
	Depth int
}

type stackItem struct {
	Data codeBlockVariableInfo
	Next *stackItem
}

type Stack struct {
	top *stackItem
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data codeBlockVariableInfo) error {
	item := &stackItem{
		Data: data,
		Next: s.top,
	}
	s.top = item
	return nil
}

func (s *Stack) Pop() (codeBlockVariableInfo, error) {
	if s.top == nil {
		return codeBlockVariableInfo{}, errors.New("stack is empty")
	}
	item := s.top
	s.top = item.Next
	return item.Data, nil
}

func (s *Stack) Peek() (codeBlockVariableInfo, error) {
	if s.top == nil {
		return codeBlockVariableInfo{}, errors.New("stack is empty")
	}
	return s.top.Data, nil
}
