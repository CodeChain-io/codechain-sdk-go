package primitives

import (
	"fmt"
	"testing"
)

func TestBigNumber(t *testing.T) {
	a := NewU128("1000112341234123413241324234")
	b := NewU128("5")
	fmt.Println(a.IsGreaterThan(&b))
	fmt.Println(a.toJSON())
}
