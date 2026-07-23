package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Printf("[log] %s\n", msg)
}

type Counter struct {
	Logger
	count int
}

func (c *Counter) Inc() {
	c.count++
	c.Log(strconv.Itoa(c.count))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var c Counter
	for i := 0; i < n; i++ {
		c.Inc()
	}
	_ = c
}
