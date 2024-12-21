package main

import (
	"testing"
)

func TestFindMinPositive(t *testing.T) {
	// Positive
	arr := []int{10, 3, 5, 8, 2}
	expected := 2
	result, err := FindMinPositive(arr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Negative
	arr = []int{-1, -2, -3}
	_, err = FindMinPositive(arr)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestSumNegative(t *testing.T) {
	// Positive
	arr := []int{-1, -3, -5, -7}
	expected := -16
	result, err := SumNegative(arr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Negative
	arr = []int{1, 2, 3}
	_, err = SumNegative(arr)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestFibonacci(t *testing.T) {
	// Positive
	n := 10
	expected := 55
	result, err := Fibonacci(n)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	// Negative
	n = -1
	_, err = Fibonacci(n)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestCalculateCurrent(t *testing.T) {
	// Positive
	voltage := 12.0
	resistance := 4.0
	expected := 3.0
	result, err := CalculateCurrent(voltage, resistance)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	// Negative
	resistance = 0.0
	_, err = CalculateCurrent(voltage, resistance)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
