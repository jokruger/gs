package unit

import (
	"context"
	"fmt"

	"github.com/jokruger/kavun"
	"github.com/jokruger/kavun/core"
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
	script := kavun.NewScript(alloc, []byte(src))

	// set values
	script.Add("a", core.IntValue(1))
	script.Add("b", core.IntValue(9))
	script.Add("c", core.IntValue(8))
	script.Add("d", core.IntValue(4))

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
