# immutable wrappers

Immutable container wrappers.

## Overview

The `immutable(x)` function wraps a container (array, bytes, dict, record, or runes) to make it immutable at the container level. Attempting to modify an immutable container raises a runtime error. Individual immutable containers can be identified via their type name (e.g., `"immutable-array"`).

## Creating Immutable Containers

### Using the `immutable()` Function

```go
a = immutable([1, 2, 3])
b = immutable(bytes("hello"))
d = immutable(dict({x: 10}))
r = immutable(u"unicode")
```

## Behavior

### Read Operations Work Normally

```go
a = immutable([1, 2, 3])
value = a[0]        // 1 (read works)
len = a.len()       // 3 (read works)
```

### Write Operations Fail

```go
a = immutable([1, 2, 3])
a[0] = 99           // runtime error - immutable
a[3] = 4            // runtime error - immutable
```

### Type Name

Immutable containers have their type names prefixed with `"immutable-"`:

```go
type_name(immutable([1, 2, 3]))         // "immutable-array"
type_name(immutable(bytes("hello")))    // "immutable-bytes"
type_name(immutable(dict({a: 1})))      // "immutable-dict"
type_name(immutable(u"text"))           // "immutable-runes"
```

## Creating Mutable Copies

The `copy()` function always returns a mutable deep copy, even from an immutable value:

```go
original = immutable([1, 2, 3])
mutable_copy = copy(original)

mutable_copy[0] = 99   // Success - copy is mutable
println(original[0])   // 1 (original unchanged)
```

## Member Functions

### `immutable()` on Containers

All immutable containers support the `immutable()` method:

**Arguments:** None

**Returns:** `immutable-container`

**Description:** Returns the same immutable container (idempotent).

```go
im_array = immutable([1, 2, 3])
im_array.immutable()    // same immutable array
```

### Other Methods

Immutable containers support the same read-only member functions as their mutable counterparts:

- **array**: `len()`, `first()`, `last()`, `min()`, `max()`, `contains()`, filtering/mapping (returns immutable results)
- **bytes**: `len()`, `first()`, `last()`, `contains()`, string conversion
- **dict**: `keys()`, `values()`, `len()`, `is_empty()`, filtering (returns immutable results)
- **runes**: `len()`, `first()`, `last()`, string conversion

```go
a = immutable([1, 2, 3, 4, 5])
filtered = a.filter(x => x > 2)
type_name(filtered)    // "immutable-array"
```

## Examples

### Protecting Data

```go
// Create protected configuration
config = immutable({
    database: "prod_db",
    timeout: 30,
    max_connections: 100
})

// Read access works
db = config.database       // "prod_db"

// Write attempt fails
config.timeout = 60        // runtime error

// Create mutable copy for modifications
mutable_config = copy(config)
mutable_config.timeout = 60  // Success
```

### Immutable Collections

```go
// Create immutable list
valid_statuses = immutable(["pending", "approved", "rejected"])

// Read operations work
if valid_statuses.contains("approved") {
    println("Valid status")
}

// Cannot accidentally modify
valid_statuses[0] = "new_status"  // runtime error
```

### API Response Protection

```go
// Simulate received API data
api_response = immutable({
    id: 123,
    user: "alice",
    data: immutable([1, 2, 3])
})

// Safe to pass around - cannot be modified
println(api_response.id)      // 123
println(api_response.data.len())  // 3

// If modification is needed:
mutable_response = copy(api_response)
```

### Function Parameters

```go
// Accept immutable parameters to signal no modification
function process_data(data) {
    // data is expected to be immutable - cannot modify
    // All operations are safe reads
    
    len = data.len()
    first = data.first()
    
    // Create mutable copy if changes needed
    result = copy(data)
    return result
}

// Pass immutable data
immutable_data = immutable([1, 2, 3, 4, 5])
result = process_data(immutable_data)
```

### Template Data

```go
// Define immutable templates
email_template = immutable({
    subject: "Welcome",
    body: "Thank you for signing up",
    from: "noreply@example.com"
})

// Create instances from template
email_instance = copy(email_template)
email_instance.to = "user@example.com"

// Original template unchanged
println(email_template.to)  // undefined
```

### Nested Immutability

```go
// Entire structure is immutable
config = immutable({
    server: {
        host: "localhost",
        port: 8080
    },
    features: immutable(["auth", "logging"])
})

// All levels are protected
config.server.host = "newhost"     // runtime error (can't modify server record)
config.features[0] = "monitoring"  // runtime error (immutable array)
```

### Defensive Copying

```go
// Receive mutable data, return immutable version
function get_safe_defaults() {
    defaults = {
        timeout: 30,
        retries: 3,
        debug: false
    }
    
    // Return immutable to prevent accidental modification
    return immutable(defaults)
}

safe_config = get_safe_defaults()
safe_config.timeout = 60  // runtime error
```

## Performance Considerations

- Immutable wrappers have minimal overhead for read operations
- `copy()` creates a deep mutable copy (not shallow)
- Immutable containers are useful for:
  - Protecting configuration data
  - Preventing accidental modifications
  - Signaling intent (this data shouldn't be modified)
  - API contracts (returned data is immutable)

## Type Checking

Use `type_name()` to identify immutable containers:

```go
function handle_data(data) {
    name = type_name(data)
    
    if name == "immutable-array" {
        println("Data is immutable array")
        // Safe to work with
    } else if name == "array" {
        println("Data is mutable array")
        // Can be modified
    }
}
```

## Notes

- Immutability applies to the container level, not to nested values
- If nested values are mutable types (arrays, dicts), they can be modified
- For complete deep immutability, ensure nested values are also wrapped
- `copy()` always produces a mutable result regardless of source mutability
- Immutable containers still support all read operations efficiently
