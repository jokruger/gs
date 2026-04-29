# bool

Boolean values representing true or false.

## Overview

Boolean values are used in control flow and logical operations. Kavun has two boolean values: `true` and `false`.

## Declaration and Usage

```go
ok = true
flag = false

// Used in control flow
if ok {
    println("ok is true")
}

// Logical operations
ok && false   // false
ok || false   // true
!ok           // false
```

## Behavior

### Logical Operations

- AND (`&&`): Returns `true` only if both operands are truthy
- OR (`||`): Returns `true` if either operand is truthy
- NOT (`!`): Inverts truthiness

```go
true && true      // true
true && false     // false
false || false    // false
true || false     // true
!true             // false
!false            // true
```

### Control Flow

Booleans are used directly in conditionals and loop conditions:

```go
if true {
    println("always runs")
}

for true {
    println("infinite loop")
    break
}

for i = 0; i < 5; i = i + 1 {
    println(i)
}
```

### Coercive Equality and Comparisons

Booleans participate in equality comparisons and can be compared with other types in limited contexts:

```go
true == true          // true
true == false         // false
true != false         // true
```

## Member Functions

### Conversion Functions

#### `bool()`
Converts to boolean.

**Arguments:** None

**Returns:** `bool`

**Description:** Returns the same boolean value.

```go
true.bool()    // true
false.bool()   // false
```

#### `int()`
Converts to integer.

**Arguments:** None

**Returns:** `int`

**Description:** Converts `true` to `1` and `false` to `0`.

```go
true.int()     // 1
false.int()    // 0

// Useful for counting true conditions
count = [true, false, true].map(b => b.int()).sum()   // 2
```

#### `string()`
Converts to string.

**Arguments:** None

**Returns:** `string`

**Description:** Converts `true` to `"true"` and `false` to `"false"`.

```go
true.string()    // "true"
false.string()   // "false"

// Used for formatting and display
message = "Status: " + ok.string()   // "Status: true"
```

## Examples

### Basic Logic

```go
// Simple boolean operations
is_valid = age >= 18 && age < 65
is_ready = !is_waiting

if is_valid && is_ready {
    println("Proceed")
}
```

### Filtering with Booleans

```go
// Using booleans in filter predicates
data = [1, 2, 3, 4, 5]
evens = data.filter(n => (n % 2) == 0)   // [2, 4]

// Boolean array operations
flags = [true, false, true, false]
count_true = flags.filter(f => f).len()   // 2
```

### Conversion for Counting

```go
// Converting booleans to integers for aggregation
test_results = [true, true, false, true]
passed = test_results.map(b => b.int()).sum()   // 3
success_rate = passed.float() / test_results.len()   // 0.75
```
