package unit

import (
	"context"
	"fmt"

	"github.com/jokruger/gs"
)

func Example() {
	src := `
each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
	sum += x
	mul *= x
})`

	// create a new Script instance
	script := gs.NewScript(alloc, []byte(src))

	// set values
	script.Add("a", alloc.NewInt(1))
	script.Add("b", alloc.NewInt(9))
	script.Add("c", alloc.NewInt(8))
	script.Add("d", alloc.NewInt(4))

	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	sum := compiled.Get("sum")
	mul := compiled.Get("mul")
	fmt.Println(sum, mul)

	// Output:
	// 22 288
}
