package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, f1, f2 := -1, 0, 1
	return func() int {
		i += 1
		if i == 0 {
			return f1
		} else if i == 1 {
			return f2
		} else {
			t := f1
			f1 = f2
			f2 += t
			return f2
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
