package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	// split into 4 chunks, goroutine each, sum total
	var wg sync.WaitGroup
	var total int
	var mu sync.Mutex
	for i := 0; i < 4; i++ {
		wg.Add(1)
		from := i * n / 4
		to := (i + 1) * n / 4
		go func(a []int) {
			defer wg.Done()
			var sum = 0
			for _, v := range a {
				sum += v
			}
			defer mu.Unlock()
			mu.Lock()
			total += sum
		}(nums[from:to])
	}
	wg.Wait()
	fmt.Println(total) // replace with the real total
}
