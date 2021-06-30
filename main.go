package main

import (
	"errors"
	"strings"
)

type node struct {
	item string
	next *node
}

type stack struct {
	top  *node
	size int
}

func (p *stack) push(name string) {
	newNode := &node{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
}

func (p *stack) pop() (string, error) {
	var item string
	if p.top == nil {
		return "", errors.New("Empty Stack")
	}
	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *stack) empty() bool {
	return p.size == 0
}

func isValid(expr string) bool {
	p := &stack{nil, 0}
	var x string
	//expr = string(expr[i])
	//expr = strings.ToLower(strings.Trim(expr, ""))
	expr = strings.Trim(expr, "")
	// Traversing the Expression
	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' || expr[i] == '[' || expr[i] == '{' {
			// Push the element in the stack
			p.push(string(expr[i]))
			continue
		}

		// IF current current character is not opening
		// bracket, then it must be closing. So stack
		// cannot be empty at this point.
		if p.empty() {
			return false
		}

		switch expr[i] {
		case ')':

			// Store the top element in a
			x, _ = p.pop()
			if x == "{" || x == "[" {
				return false
			}

		case '}':

			// Store the top element in b
			x, _ = p.pop()
			if x == "(" || x == "[" {
				return false
			}

		case ']':
			x, _ = p.pop()
			if x == "(" || x == "{" {
				return false
			}
		}
	}
	// Check Empty Stack
	return (p.empty())
}
