package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	c := coord{1, 2, 3}
	for rot := 0; rot < 24; rot++ {
		rc := c.rotate(rot)
		fmt.Printf("%d: %v\n", rot, rc)
	}
}
