package main
import "fmt"

func main() {
    var a [5]int = [5]int{1, 2, 3, 4, 5}
	n := len(a) - 1
	c := a[n]
	a[n]= a[0]
	a[0] = c

	fmt.Print(a)
}
