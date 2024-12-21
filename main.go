package main

import (
	"errors"
)

func FindMinPositive(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array is empty")
	}
	min := int(^uint(0) >> 1) // max_int
	for _, num := range arr {
		if num > 0 && num < min {
			min = num
		}
	}
	if min == int(^uint(0)>>1) {
		return 0, errors.New("no positive numbers found")
	}
	return min, nil
}

func SumNegative(arr []int) (int, error) {
	sum := 0
	foundNegative := false
	for _, num := range arr {
		if num < 0 {
			sum += num
			foundNegative = true
		}
	}
	if !foundNegative {
		return 0, errors.New("no negative numbers found")
	}
	return sum, nil
}

func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev = curr
		curr = next
	}
	return curr, nil
}

func CalculateCurrent(voltage, resistance float64) (float64, error) {
	if resistance == 0 {
		return 0, errors.New("resistance cannot be zero")
	}
	return voltage / resistance, nil
}

func main() {
	
}
