# CPU Temperature Monitor

Simple Go library to get CPU temperature.  
Windows: via WMI ACPI.  
Linux: via /sys/class/thermal.  
## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/northwindlight/cputemp"
)

func main() {
    temp, err := cputemp.GetCPUTemperature()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("CPU Temperature: %.1f°C\n", temp)
}
