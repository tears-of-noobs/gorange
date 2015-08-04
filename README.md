# gorange

Make range (slice) of values between first and last value

### Install package
```
go get github.com/tears-of-noobs/gorange
```

### Use it

```go

package main                                                                                                                          
                                                                                                                               
import (                                                                                                                              
    "fmt"                                                                                                                             
    "gorange"                                                                                                                         
)                                                                                                                                     
                                                                                                                                      
func main() {                                                                                                                         
                                                                                                                                      
    rn, err := gorange.IntRange(1, 5)                                                                                               
    if err != nil {                                                                                                                   
        fmt.Println(err.Error())                                                                                                      
    }                                                                                                                                 
    fmt.Printf("%v\n", rn)                                                                                                            
                                                                                                                                      
} 
// Output - [1 2 3 4 5]
```
