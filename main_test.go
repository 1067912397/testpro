package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expect := 3
	actually := Add(1, 2)
	if expect != actually {
		fmt.Errorf("Test error!")
	}
}