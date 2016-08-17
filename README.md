# loglite

A light weight logger for golang. Two loggers for Debug and Info.

# Example

```go
package main

import (
    "github.com/shinmyung0/loglite"
)

func main() {

    loglite.Info("Running program")
    loglite.Debug("Passed in as arguments", os.Args)


}


```

Example output

```
2016/08/17 20:39:03 loglite_test.go:11 [INFO] Running program
2016/08/17 20:39:03 loglite_test.go:12 [DEBUG] Passed in as arguments []  

```
