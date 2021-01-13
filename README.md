# Errors

```
go get github.com/a-novel/errors
```

A simple package for better errors handling. It assigns a string ID to each error for better parsing. It is also
duck typed with Go error for interoperability.

```go
package myPackage

import
import (
	"fmt"
	"github.com/a-novel/errors"
)

func main() {
    err := errors.New("uniq_id", "something happened")
    
    // ... 
    
    fmt.Print(err.Error()) // something happened
    fmt.Print(err.ID == "uniq_id") // true
    
    switch err.ID {
    case "uniq_id":
        // Do something.
    case "uniq_id_2":
        // Do something else.
    default:
        // Unknown/Unhandled error.
    }
}
```

# License

Distributed under [Apache License 2.0](https://www.github.com/a-novel/errors/blob/master/LICENSE)