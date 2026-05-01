# error

First-class error values.

## Overview

The `error` type represents an error or exceptional condition. Errors are first-class values in Kavun, meaning they can be stored, passed around, and operated on like any other value. This allows for elegant error handling and propagation.

## Declaration and Creation

### Construction

```go
e = error("something went wrong")
e2 = error("Database connection failed")
e3 = error("Invalid input: expected integer")
```

### From Values

```go
message = "Network timeout"
err = error(message)
```

## Member Functions

### Accessor Functions

#### `value()`

Gets the error message.

**Arguments:** None

**Returns:** `any`

**Description:** Returns the payload/message that was wrapped in the error.

```go
e = error("something went wrong")
e.value()    // "something went wrong"

// Error with complex payload
details = {code: 404, message: "Not found"}
e = error(details)
e.value()    // {code: 404, message: "Not found"}
```

### Conversion Functions

#### `string()`

Converts to string.

**Arguments:** None

**Returns:** `string`

**Description:** Returns the error message as a string. If the error payload is not a string, it attempts to convert it to string format.

```go
e = error("something went wrong")
e.string()   // "something went wrong"

// Error with non-string payload
e2 = error(404)
e2.string()  // "404"
```

## Built-in Error Functions

### Error Detection

#### `is_error(x)`

Checks if a value is an error.

**Arguments:**

- `x` (any): Value to check

**Returns:** `bool`

**Description:** Returns `true` if the value is an error, `false` otherwise.

```go
e = error("failed")
is_error(e)           // true

value = 42
is_error(value)       // false

undefined_val = undefined
is_error(undefined_val)  // false
```

## Examples

### Basic Error Handling

```go
fmt = import("fmt")

// Create and check errors
result = error("operation failed")

if is_error(result) {
    fmt.println("Error occurred: " + result.string())
}
```

### Error Propagation

```go
fmt = import("fmt")

// Function that returns error on failure
function divide(a, b) {
    if b == 0 {
        return error("division by zero")
    }
    return a / b
}

result = divide(10, 0)
if is_error(result) {
    fmt.println("Calculation failed: " + result.value().string())
}
```

### Error with Structured Data

```go
fmt = import("fmt")

// Error with detailed information
function validate_user(data) {
    if data.name == undefined || data.name == "" {
        return error({
            code: "INVALID_NAME",
            message: "Name is required",
            field: "name"
        })
    }

    if data.age == undefined || data.age < 0 {
        return error({
            code: "INVALID_AGE",
            message: "Age must be non-negative",
            field: "age"
        })
    }

    return data
}

user = {name: "", age: 25}
result = validate_user(user)

if is_error(result) {
    details = result.value()
    fmt.println("Validation failed")
    fmt.println("Code: " + details.code)
    fmt.println("Field: " + details.field)
}
```

### Error Aggregation

```go
fmt = import("fmt")

// Collect multiple errors
function validate_form(form) {
    errors = []

    if form.email == undefined || form.email == "" {
        errors = errors + [error("Email is required")]
    }

    if form.password == undefined || form.password.len() < 8 {
        errors = errors + [error("Password must be at least 8 characters")]
    }

    if form.age != undefined && form.age < 18 {
        errors = errors + [error("Must be 18 or older")]
    }

    if errors.len() > 0 {
        return error({
            message: "Multiple validation errors",
            count: errors.len(),
            errors: errors
        })
    }

    return form
}

form = {email: "", password: "short", age: 16}
result = validate_form(form)

if is_error(result) {
    details = result.value()
    fmt.println("Found " + details.count.string() + " errors")
}
```

### Conditional Error Handling

```go
// Handle errors conditionally
function safe_int(value) {
    result = value.int()
    if result == undefined {
        return error("Cannot convert to integer: " + value.string())
    }
    return result
}

// Use with fallback
function parse_with_default(value, default) {
    result = safe_int(value)
    if is_error(result) {
        return default
    }
    return result
}

port = parse_with_default("8080", 3000)      // 8080
port = parse_with_default("invalid", 3000)   // 3000
```

### Error Filtering

```go
fmt = import("fmt")

// Filter operations
data = [1, 2, "three", 4, "five", 6]

// Convert with error handling
converted = []
for item in data {
    result = item.int()
    if !is_error(result) && result != undefined {
        converted = converted + [result]
    }
}

fmt.println(converted)  // [1, 2, 4, 6]
```

## Error Value Payloads

Errors can wrap any type of value:

- **String**: Simple error messages
- **Record/Dict**: Structured error information with code, message, context
- **Integer**: Error codes
- **Array**: Multiple error details
- **Any other type**: Custom error representations

```go
simple = error("failed")
structured = error({code: "ERR_001", msg: "Failed"})
numeric = error(500)
details = error([error("Subissue 1"), error("Subissue 2")])
```

## Common Patterns

### Early Return on Error

```go
function process_data(data) {
    validated = validate(data)
    if is_error(validated) {
        return validated  // Propagate error
    }

    transformed = transform(validated)
    if is_error(transformed) {
        return transformed  // Propagate error
    }

    return transformed
}
```

### Default Values with Errors

```go
// Use error values in computations by checking first
value = get_config_value("timeout")
timeout = if is_error(value) { 30 } else { value }
```

### Logging Errors

```go
fmt = import("fmt")

function log_error(err) {
    if is_error(err) {
        message = err.string()
        payload = err.value()
        fmt.println("[ERROR] " + message)
        // In real code, write to log file or service
    }
}

log_error(error("Something went wrong"))
```

## Design Notes

- Errors are values, not exceptions - they don't interrupt execution
- Use conditional checks with `is_error()` to handle errors
- Errors can be returned from functions or stored in data structures
- The payload of an error can be any type, allowing flexible error representation
- Functions should return `error` values on failure rather than throwing exceptions
