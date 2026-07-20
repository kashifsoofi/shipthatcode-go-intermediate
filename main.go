package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a := make(chan int)
	b := make(chan int)
	go func() {
		defer close(a)
		for _, f := range fields {
			n, _ := strconv.Atoi(f)
			// TODO: send n into channel a
			a <- n
		}
	}()
	go func() {
		defer close(b)
		for n := range a {
			// TODO: send the square of n into channel b
			b <- n * n
		}
	}()
	sum := 0
	for v := range b {
		sum += v
	}
	fmt.Println(sum)
}
