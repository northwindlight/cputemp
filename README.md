# CPU Temperature Monitor

Simple Go library to get CPU temperature via WMI ACPI.

## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/yourusername/cputemp"
)

func main() {
    temp, err := cputemp.GetCPUTemperature()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("CPU Temperature: %.1f°C\n", temp)
}
