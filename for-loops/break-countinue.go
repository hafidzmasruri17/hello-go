package main

import "fmt"

func main() {

	for i := 0; i <= 20; i++ {

		if i%2 == 1 {
			continue
		}

		if i > 16 {
			break
		}

		fmt.Println(i)

	}

}