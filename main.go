package main

import (
	"fmt"

	"postsort/postsort"
)

func main() {
	res, err := postsort.Sort(10, 10, 10, 10)
	fmt.Print(res, err)
}
