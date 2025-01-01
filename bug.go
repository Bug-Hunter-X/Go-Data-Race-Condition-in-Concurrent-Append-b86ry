```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var data []int
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        data = append(data, i)
                }(i)
        }
        wg.Wait()
        fmt.Println(len(data))
}
```