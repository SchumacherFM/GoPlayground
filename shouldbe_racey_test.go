package GoPlayground

import (
	"fmt"
	"testing"
)

func TestShouldBeRacy(t *testing.T) {
	a := new(int)
	*a = 1
	go func() {
		for i := 0; i < 1e5; i++ {
			*a = i
		}
		println("gorutine terminated")
	}()
	var x *int
	//var y int
	for i := 0; i < 1e5; i++ {
		x = a
		//y = *a
	}
	fmt.Printf("p:%p\n", x)
}
