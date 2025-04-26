package main
import "fmt"

func main() {
    var a []int = make([]int , 5)

	mx := a[0]
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
		if mx < a[i] {
			mx = a[i]
		}
	}
	fmt.Printf("max = %d\n", mx)
}
