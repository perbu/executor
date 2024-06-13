# executor

A Go package to execute shell scripts in a temporary file and return stdout, stderr and exit code.

## Installation

```
go get github.com/perbu/executor
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/perbu/executor"
)

func main() {
    script := []byte("#!/bin/sh\necho 'Hello, World!'\n")
    err := executor.Execute(script)
    if err != nil {
        fmt.Printf("Error executing script: %v\n", err)
        if execErr, ok := err.(*executor.ErrExecute); ok {
            fmt.Printf("Stdout: %s\n", execErr.Stdout)
            fmt.Printf("Stderr: %s\n", execErr.Stderr)
            fmt.Printf("Exit code: %d\n", execErr.ExitCode)
        }
    }
}
```

## Error Handling

If the script fails to execute, `Execute` will return an `ErrExecute` error which contains:
- `SubError`: The underlying error that caused the execution to fail
- `Stdout`: The contents of stdout 
- `Stderr`: The contents of stderr
- `ExitCode`: The exit code of the script

You can type assert the returned error to `ErrExecute` to access these fields:

```go
if execErr, ok := err.(*executor.ErrExecute); ok {
    fmt.Printf("Stdout: %s\n", execErr.Stdout) 
    fmt.Printf("Stderr: %s\n", execErr.Stderr)
    fmt.Printf("Exit code: %d\n", execErr.ExitCode)
}
```

## Testing

The package includes unit tests covering success and error scenarios. To run the tests:

```
go test ./...
```

## License 

See LICENSE.md

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
