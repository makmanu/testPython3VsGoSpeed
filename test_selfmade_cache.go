package main

import(
	"fmt"
	"time"
)

func fib_cache() func(int) int{
	cache := map[int]int{}

	var fib func(n int) int
	fib = func(n int) int{
		if n < 2{
			return n
		}
		if val, ok := cache[n]; ok {
			return val
		}
		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}
	return fib
}


func main(){
	timeStart := time.Now()
    for i := 0; i < 45; i++ {
        fmt.Println(fib_cache()(i))
		fmt.Println("		", time.Since(timeStart))
		if time.Since(timeStart).Seconds() > 10{
			return
		}
	}
	fmt.Println(time.Since(timeStart))
}