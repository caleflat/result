# A result type for Golang 1.19+
This was a small project to help handle errors in a more functional way. It is inspired by Rust's Result type.

# Usage
```go
package main

import (
    "errors"
    "fmt"
    "github.com/caleflat/result"
)

func main() {
    res := divide(10, 2)
    fmt.Println(res) // Ok(5)
    
    if res.IsOk() {
        fmt.Println(res.Unwrap()) // 5
    }
    
    res = divide(10, 0)
    fmt.Println(res) // Err("divide by zero")
    
    if res.IsErr() {
        fmt.Println(res.UnwrapErr()) // divide by zero
    }
}

func divide(a, b int) result.Result {
    if b == 0 {
        return result.Err(errors.New("cannot divide by zero"))
    }
    return result.Ok(a / b)
}