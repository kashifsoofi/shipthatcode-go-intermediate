package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 0

	// TODO: split `nums` into 4 chunks and launch a goroutine per chunk.
	// Each goroutine should sum its chunk and add the result into `total`,
	// guarded by `mu`. Use `wg` (Add/Done/Wait) to wait for all goroutines
	// to finish before printing.
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

	fmt.Println(total)
}
