# gorange

Make range (slice) of values between first and last value with different step (like python range() function)

### Install package
```
go get github.com/tears-of-noobs/gorange
```

### Use it

```go

package main                                                                                                                          
                                                                                                                               
import (                                                                                                                              
    "fmt"                                                                                                                             

    "github.com/tears-of-noobs/gorange"                                                                                                                         
)                                                                                                                                     
                                                                                                                                      
func main() {                                                                                                                         
                                                                                                                                      
    rn, err := gorange.IntRange(1, 10, 2)
    if err != nil {                                                                                                                   
        fmt.Println(err.Error())                                                                                                      
    }                                                                                                                                 
    fmt.Printf("%v\n", rn)                                                                                                            
                                                                                                                                      
} 
// Output - [1 3 5 7 9]
```
