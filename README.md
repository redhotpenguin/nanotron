# nanotron

## takes JSON trace data, transforms it into protobuf or other format
## currently setup for [tracemate](https://github.com/MajorLeagueBaseball/tracemate/) traces

```
package main
import (
    tm "github.com/redhotpenguin/nanotron/tracemate'
    "io/ioutil"
    "fmt"
)

func main() {
    ct, _ := ioutil.Readfile("tracemate.json")
    proto := tm.JsonToProto(ct)
    fmt.Println("proto conversion is ", proto) 
}
```
