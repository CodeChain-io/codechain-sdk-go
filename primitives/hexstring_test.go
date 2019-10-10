package primitives

import (
	"fmt"
	"testing"
)

func TestH128(t *testing.T) {
	a := NewH128Zero()
	b := NewH128Zero()
	c := H128([16]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	d := NewH512Zero()
	fmt.Println(d)
	fmt.Println(d.ToJSON())
	fmt.Println(a.Cmp(b))
	fmt.Println(c.Cmp(b))
	fmt.Println(a.ToString())
	fmt.Println(a.RlpBytes())
	fmt.Println(a.ToJSON())
}
