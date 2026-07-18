package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func validateAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}

	if age < 0 {
		return 0, errors.New("negative")
	}

	return age, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, err := validateAge(sc.Text())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", age)
	}
}
