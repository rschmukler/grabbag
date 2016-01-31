# GrabBag

Simple nested object traversal for Go

## Install

```
go get github.com/rschmukler/grabbag
```

## Examples

```go
import (
  "github.com/rschmukler/grabbag"
)

var data := map[string]interface{}{
  "number": 5,
  "string": "Hello world",
  "nested": map[string]interface{}{
    "bool": true,
  },
}

func main() {
  gb := grabbag.FromData(data)

  num := gb.Int("number")
  b := gb.Bool("nested.bool")
}
```
