package main

import (
	"strings"
)

type calc struct {
	stack []float64
}

func (c *calc) isCalculable() bool {
	if len(c.stack) <= 1 {
		return false
	} else {
		return true
	}
}

func (c *calc) addStack(input float64) {
	c.stack = append(c.stack, input)
}

func (c *calc) delStack() {
	c.stack = c.stack[:len(c.stack) - 2]
}

func (c *calc) calculation(input string) float64 {
	a := c.stack[len(c.stack) - 2]
	b := c.stack[len(c.stack) - 1]
	switch {
	case strings.Contains(input, "+"):
		return a + b
	case strings.Contains(input, "-"):
		return a - b
	case strings.Contains(input, "*"):
		return a * b
	case strings.Contains(input, "/"):
		if b == 0 {
			return 0
		} else {
			return a / b
		}
	default:
		return 0
	}
}

