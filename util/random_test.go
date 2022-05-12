package util

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRandomInt64(t *testing.T) {
	n := RandomInt64()

	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("int 64:", i)
}

func TestRandomInt32(t *testing.T) {
	n := RandomInt32()

	i, err := strconv.ParseInt(n, 10, 32)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("int 32:", i)
}
