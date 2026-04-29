# bytes

Mutable byte arrays.

## Overview

The `bytes` type represents a mutable sequence of bytes (integers 0-255). Use `bytes` when you need to manipulate raw byte data or when mutability is required. Each index holds a byte value (0 to 255).

## Declaration and Usage

### Construction

```go
b = bytes("abc")              // from string
b2 = [97, 98, 99].bytes()     // from array
empty = bytes()               // empty bytes
```

### Indexing and Slicing

```go
b = bytes("abc")
b[0]                          // 97 (first byte, as int)
b[0:2]                        // bytes slice
```

### Mutability

```go
b = bytes("hello")
b[0] = 72                     // Change first byte to 'H'
b.string()                    // "Hello"

// Concatenation
b1 = bytes("ab")
b2 = bytes("cd")
result = b1 + b2              // bytes with [97, 98, 99, 100]
```

## Member Functions

### Conversion Functions

#### `bytes()`
Converts to bytes.

**Arguments:** None

**Returns:** `bytes`

**Description:** Returns the same bytes value.

```go
bytes("hello").bytes()    // bytes("hello")
```

#### `array()`
Converts to array of integers.

**Arguments:** None

**Returns:** `array`

**Description:** Returns an array of integers (0-255) representing the bytes.

```go
bytes("ABC").array()      // [65, 66, 67]
```

#### `string()`
Converts to string.

**Arguments:** None

**Returns:** `string`

**Description:** Interprets the bytes as UTF-8 and returns a string. May return invalid UTF-8 as-is.

```go
bytes("hello").string()   // "hello"
[72, 105].bytes().string()  // "Hi"
```

#### `record()`
Converts to record.

**Arguments:** None

**Returns:** `record`

**Description:** Converts bytes to a record where keys are string indices (`"0"`, `"1"`, ...), and values are byte values as ints.

```go
bytes("abc").record()   // {"0": 97, "1": 98, "2": 99}
```

#### `dict()`
Converts to dict.

**Arguments:** None

**Returns:** `dict`

**Description:** Converts bytes to a dict where keys are string indices (`"0"`, `"1"`, ...), and values are byte values as ints.

```go
bytes("abc").dict()      // dict({"0": 97, "1": 98, "2": 99})
```

### Transformation and Filtering Functions

#### `sort()`
Sorts bytes in ascending order.

**Arguments:** None

**Returns:** `bytes`

**Description:** Returns a new bytes with values sorted from smallest to largest.

```go
bytes("dcba").sort()     // bytes("abcd")
bytes([3, 1, 4, 1]).sort()  // bytes([1, 1, 3, 4])
```

#### `filter(fn)`
Filters by predicate.

**Arguments:**
- `fn` (function): Predicate that takes one argument `(byte)` or two arguments `(index, byte)` and returns bool

**Returns:** `bytes`

**Description:** Returns bytes containing only values where the predicate returns `true`.

```go
bytes("hello123").filter(b => b >= 'a'.int() && b <= 'z'.int())  
// bytes("hello")

bytes([1, 2, 3, 4, 5]).filter(b => b % 2 == 0)  // bytes([2, 4])
```

### Predicate Functions

#### `all(fn)`
Tests if all bytes match predicate.

**Arguments:**
- `fn` (function): Predicate that takes one argument `(byte)` or two arguments `(index, byte)` and returns bool

**Returns:** `bool`

**Description:** Returns `true` if all bytes satisfy the predicate.

```go
bytes("abc").all(b => b >= 'a'.int() && b <= 'z'.int())   // true
bytes("abc123").all(b => b >= 'a'.int() && b <= 'z'.int()) // false
```

#### `any(fn)`
Tests if any byte matches predicate.

**Arguments:**
- `fn` (function): Predicate that takes one argument `(byte)` or two arguments `(index, byte)` and returns bool

**Returns:** `bool`

**Description:** Returns `true` if any byte satisfies the predicate.

```go
bytes("abc").any(b => b >= '0'.int() && b <= '9'.int())      // false
bytes("abc123").any(b => b >= '0'.int() && b <= '9'.int())   // true
```

### Aggregation Functions

#### `count(fn)`
Counts bytes matching predicate.

**Arguments:**
- `fn` (function): Predicate that takes one argument `(byte)` or two arguments `(index, byte)` and returns bool

**Returns:** `int`

**Description:** Returns the number of bytes where the predicate returns `true`.

```go
bytes("hello world").count(b => b == ' '.int())    // 1
bytes("a0b1c2").count(b => b >= '0'.int() && b <= '9'.int())  // 3
```

#### `min()`
Finds minimum byte.

**Arguments:** None

**Returns:** `int | undefined`

**Description:** Returns the smallest byte value (0-255). Returns `undefined` for empty bytes.

```go
bytes("hello").min()    // 101 ('e')
bytes().min()           // undefined
```

#### `max()`
Finds maximum byte.

**Arguments:** None

**Returns:** `int | undefined`

**Description:** Returns the largest byte value (0-255). Returns `undefined` for empty bytes.

```go
bytes("hello").max()    // 111 ('o')
bytes().max()           // undefined
```

### Query and Accessor Functions

#### `is_empty()`
Checks if bytes is empty.

**Arguments:** None

**Returns:** `bool`

**Description:** Returns `true` if the bytes has zero bytes.

```go
bytes().is_empty()      // true
bytes("hello").is_empty() // false
```

#### `len()`
Gets byte count.

**Arguments:** None

**Returns:** `int`

**Description:** Returns the number of bytes.

```go
bytes("hello").len()    // 5
bytes([1, 2, 3]).len()  // 3
```

#### `first()`
Gets first byte.

**Arguments:** None

**Returns:** `int | undefined`

**Description:** Returns the first byte as an integer (0-255). Returns `undefined` for empty bytes.

```go
bytes("hello").first()  // 104 ('h')
bytes().first()         // undefined
```

#### `last()`
Gets last byte.

**Arguments:** None

**Returns:** `int | undefined`

**Description:** Returns the last byte as an integer (0-255). Returns `undefined` for empty bytes.

```go
bytes("hello").last()   // 111 ('o')
bytes().last()          // undefined
```

#### `contains(x)`
Checks if bytes contains a value.

**Arguments:**
- `x` (int): Byte value to search for (0-255)

**Returns:** `bool`

**Description:** Returns `true` if the byte value is found.

```go
bytes("hello").contains('h'.int())    // true
bytes("hello").contains('x'.int())    // false
bytes([1, 2, 3]).contains(2)          // true
```

## Examples

### Binary Data Manipulation

```go
// Create and modify binary data
data = [0xFF, 0x00, 0x42]
data[1] = 0xAA           // Modify a byte
println(data.string())   // Print as string (may be non-printable)
```

### String Encoding/Decoding

```go
// Convert string to bytes and back
original = "Hello"
binary = original.bytes()  // Convert to bytes

// Modify
binary[0] = 'J'.int()      // Change 'H' to 'J'

result = binary.string()   // "Jello"
```

### Byte Filtering and Analysis

```go
// Filter ASCII text
text = bytes("Hello123!")
letters = text.filter(b => 
    (b >= 'A'.int() && b <= 'Z'.int()) ||
    (b >= 'a'.int() && b <= 'z'.int())
)
println(letters.string())   // "Hello"

// Extract digits
digits = text.filter(b => b >= '0'.int() && b <= '9'.int())
println(digits.string())    // "123"
```

### Data Statistics

```go
// Analyze byte distribution
data = bytes("programming")

min_byte = data.min()       // 97 ('a')
max_byte = data.max()       // 114 ('r')
total_bytes = data.len()    // 11

// Count specific bytes
letter_a_count = data.count(b => b == 'a'.int())  // 1
```

### JSON Processing

```go
// Simulate JSON manipulation
json_bytes = `{"name":"Alice","age":30}`.bytes()

// Convert to string for processing
json_str = json_bytes.string()
record = json_str.record()

// Verify data integrity
if record != undefined {
    println("Valid JSON")
}
```

## Performance Notes

- `bytes` is mutable, so modifications happen in-place
- `bytes` maintains reference semantics: `a = b` makes both variables point to the same bytes
- Use `copy()` to create independent copies
- Byte values must be in range 0-255; values outside this range raise errors
