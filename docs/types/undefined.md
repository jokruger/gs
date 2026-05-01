# undefined

Represents the absence of a value.

## Overview

The `undefined` value is used to represent the absence of a meaningful value. It's returned in situations where:

- A field or index doesn't exist
- A conversion fails (unless a fallback is provided)
- Operations attempt to access non-existent resources

## Behavior

### Field and Index Access

Any field or index access on `undefined` returns `undefined`:

```go
u = undefined
u.anything        // undefined
u[0]              // undefined
u.deeply.nested   // undefined
```

### Truthiness

`undefined` is falsy in boolean contexts:

```go
if undefined {
    // This block is NOT executed
}

undefined && true   // false
undefined || false  // false
```

### Conversion Fallbacks

Many conversion builtins return `undefined` on conversion failure unless a fallback is provided:

```go
int("not a number")           // undefined
int("not a number", 0)        // 0 (uses fallback)

float("invalid")              // undefined
float("invalid", 3.14)        // 3.14 (uses fallback)
```

## Member Functions

`undefined` has no member functions. Attempting to call a method on `undefined` will result in a runtime error.
