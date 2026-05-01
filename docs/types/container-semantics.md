# container semantics

Shared semantics of array-like containers (`array`, `bytes`, `runes`, `dict`, `record`): reference behavior,
immutability, propagation through derived operations, and the aliasing pitfalls of `append`.

## Reference Semantics

Containers are reference-typed. Assignment shares the underlying buffer; mutation through one variable is visible
through any other variable that refers to the same container:

```go
fmt = import("fmt")
a = [1, 2, 3]
b = a
b[0] = 99
fmt.println(a)              // [99, 2, 3] - a sees b's mutation
```

Use `copy()` to obtain an independent value:

```go
fmt = import("fmt")
a = [1, 2, 3]
b = copy(a)
b[0] = 99
fmt.println(a)              // [1, 2, 3] - unchanged
```

## Immutable Wrappers

The `immutable(x)` function wraps a container (`array`, `bytes`, `dict`, `record`, `runes`) to make it immutable at the
container level. Attempting to modify an immutable container raises a runtime error. Individual immutable containers
can be identified via their type name (e.g., `"immutable-array"`).

### Creating Immutable Containers

```go
a = immutable([1, 2, 3])
r = immutable({x: 10})
```

### Read Operations Work Normally

```go
a = immutable([1, 2, 3])
value = a[0]            // 1 (read works)
len = a.len()           // 3 (read works)
```

### Write Operations Fail

```go
a = immutable([1, 2, 3])
a[0] = 99               // runtime error - immutable
a[3] = 4                // runtime error - immutable
```

### Type Name

Immutable containers have their type names prefixed with `"immutable-"`:

```go
type_name(immutable([1, 2, 3]))         // "immutable-array"
type_name(immutable({a: 1}))            // "immutable-record"
type_name(immutable(dict({a: 1})))      // "immutable-dict"
type_name(immutable(bytes("ab")))       // "immutable-bytes"
type_name(immutable(runes("ab")))       // "immutable-runes"
```

### Creating Mutable Copies

The `copy()` function always returns a mutable deep copy, even from an immutable value:

```go
fmt = import("fmt")
original = immutable([1, 2, 3])
mutable_copy = copy(original)

mutable_copy[0] = 99    // Success - copy is mutable
fmt.println(original[0])    // 1 (original unchanged)
```

## Propagation Through Slicing and Chunking

Whether the immutable flag carries over to a derived value depends on whether that value shares memory with the source:

- **Two-part slice** (`v[a:b]`): the result is a view over the same backing buffer, so it inherits the source's
  immutability. Slicing an `immutable-array` / `immutable-bytes` / `immutable-runes` yields another immutable value.
- **Stepped slice** (`v[a:b:s]`): always builds a fresh independent buffer, so the result is always mutable.
- **`chunk(n)`** (default): each chunk is a view over the source buffer, so chunks inherit the source's immutability.
- **`chunk(n, true)`**: each chunk is an independent copy, so chunks are always mutable.
- **`copy()`, `sort()`, `reverse()`, `filter()`, `map()`**: always return a fresh mutable value.

```go
a = immutable([1, 2, 3, 4])
type_name(a[1:3])               // "immutable-array"  (shares memory)
type_name(a[::-1])              // "array"             (fresh buffer)
type_name(a.chunk(2)[0])        // "immutable-array"
type_name(a.chunk(2, true)[0])  // "array"
```

## Append Aliasing

`append` returns a value that **may or may not** share its backing buffer with the source, depending on the source's
spare capacity. This means assigning the result of `append` to a variable other than the source produces unpredictable
behavior.

### The Safe Pattern

Always assign the result of `append` back to the source variable:

```go
x = [1, 2, 3]
x = append(x, 4)        // safe - reassign to x
x = append(x, 5, 6)     // safe
```

In this pattern, the source variable is the only handle to the (possibly relocated) buffer, so aliasing cannot cause
surprises.

### The Unsafe Pattern

Storing `append`'s result in a different variable while keeping the source around exposes implementation-defined
aliasing:

```go
fmt = import("fmt")
x = [1, 2, 3]
v1 = append(x, 100)     // v1 = [1, 2, 3, 100]
v2 = append(x, 200)     // v2 = [1, 2, 3, 200]
fmt.println(v1)         // ??? could be [1, 2, 3, 100] OR [1, 2, 3, 200]
```

The outcome depends on the hidden capacity of `x`'s backing buffer:

- If `x` has spare capacity, both appends write into the same memory at the same offset. `v2`'s write **overwrites**
  the slot `v1` exposes, so `v1` and `v2` end up equal — both showing `200`.
- If `x` has no spare capacity, the first append allocates a new buffer for `v1`; the second append again allocates,
  producing an independent `v2`. `v1` keeps `100`.

You cannot rely on either outcome — the capacity is an internal detail that may change with the size of the container,
the runtime allocator, or future versions of Kavun.

### How to Get Predictable Behavior

If you need an independent extended container without disturbing the source, copy first:

```go
x = [1, 2, 3]
v1 = append(copy(x), 100)   // v1 is independent of x
v2 = append(copy(x), 200)   // v2 is independent of x and v1
```

The same rules apply to `bytes` and `runes`. `append` on an immutable container is rejected at runtime, so the aliasing
pitfall only applies to mutable sources.

## Notes

- Immutability applies to the container level, not to nested values.
- If nested values are mutable types (arrays, dicts), they can still be modified through any reference to them.
- For complete deep immutability, ensure nested values are also wrapped.
- `copy()` always produces a mutable result regardless of source mutability.
- Immutable containers still support all read operations efficiently.
