package benchmark

import (
	"testing"

	"github.com/jokruger/kavun"
)

func BenchmarkVM(b *testing.B) {
	//src := []byte(`out = range(1, 10000, 1).to_array().reduce(0, (a, b) => a + b * b)`)

	src := []byte(`
out = decimal(0)
for i := 0; i < 1000; i++ {
	out = out + decimal(i)
}
`)

	script := kavun.NewScript(src)
	compiled, err := script.Compile(nil, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("vmRun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := compiled.Run(); err != nil {
				b.Fatal(err)
			}
		}
	})
}
