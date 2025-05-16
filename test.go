package main

import(
	"fmt"
	"time"
)

func fib(n int) int{
    if n < 2{
        return n
	}
    return fib(n-1) + fib(n-2)
}

func main(){
	timeStart := time.Now()
    for i := 0; i < 45; i++ {
		fmt.Printf("fib number №%d\n", i)		
        fmt.Println(fib(i))
		fmt.Println("		", time.Since(timeStart))
		if time.Since(timeStart).Seconds() > 10{
			return
		}
	}
	fmt.Println(time.Since(timeStart))
}