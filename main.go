package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) > 0 {
		len := len(s.items)
		item := s.items[len-1]
		s.items = s.items[:len-1]
		return item, true
	}

	return 0, false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	parts := strings.Fields(sc.Text())
	var s Stack
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		_ = n
		s.Push(n)
	}
	for {
		x, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(x)
	}
	_ = fmt.Print
}
