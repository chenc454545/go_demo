package main
  
import (
        "fmt"
)

func main() {
        x := 75
        y := 50
        fmt.Println("gcd", gcd(x,y))
        fmt.Println("fib", fib(x))

	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)
}

func gcd(x, y int) int {
        fmt.Printf("gcd start:[%d][%d]\n", x, y)
	for y != 0 {
                x, y = y, x%y
	}
        return x
}

func fib(n int) int {
        x,y := 0,1
        for i:= 0; i<n; i++ {
                x,y = y, x+y
        }
        return x
}

