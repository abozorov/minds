package main
import "fmt"

func main() {
    var a []int = make([]int , 0, 10)

	for i := 1; i <= 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	var b []int
	b = a[1:10:2]
	fmt.Println(b)
}
