# loglite

A light weight logger for golang. A global logger initialized for your entire app.
Two log levels : Debug and Log. Debug is for verbose information helpful for debugging (Don't go overboard though).
Log is for information useful to the user.

# Example

```go
package main

import (
    "github.com/shinmyung0/loglite"
)

func main() {

    // loglite already has a global logger initialized
    loglite.Logln("Running program")
    loglite.Debugln("Passed in as arguments", os.Args)


}


```

Example output

```
2016/08/17 20:39:03 loglite_test.go:11 [LOG]  
2016/08/17 20:39:03 loglite_test.go:11 [DEBUG]  

```
