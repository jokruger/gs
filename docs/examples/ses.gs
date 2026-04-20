#!/usr/bin/env gs

fmt = import("fmt")

result = [1, 2, 3, 4, 5, 6]
  .filter(x => x % 2 == 0)
  .map(x => x * x)
  .reduce(0, (sum, x) => sum + x)

fmt.printf("sum of even squares: %v\n", result)