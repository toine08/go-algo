package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/toine08/go-algo/algorithm"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{3, 0, -1, 8, 7, 2}, expected: []int{-1, 0, 2, 3, 7, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: generateLargeArray(10000), expected: generateSortedLargeArray(10000)},
	}

	for _, test := range tests {
		algorithm.Insertion_sort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("InsertionSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}

func TestInsertionSortConcurrent(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{3, 0, -1, 8, 7, 2}, expected: []int{-1, 0, 2, 3, 7, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: generateLargeArray(10000), expected: generateSortedLargeArray(10000)},
	}

	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		start := time.Now()
		done := make(chan bool)
		go func() {
			algorithm.Insertion_sort(inputCopy)
			done <- true
		}()

		<-done
		duration := time.Since(start)

		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("InsertionSortConcurrent(%v) = %v; want %v", test.input, inputCopy, test.expected)
		}

		fmt.Printf("InsertionSortConcurrent(%v) took %v\n", test.input, duration)
	}
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{3, 0, -1, 8, 7, 2}, expected: []int{-1, 0, 2, 3, 7, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: generateLargeArray(10000), expected: generateSortedLargeArray(10000)},
	}

	for _, test := range tests {
		algorithm.Bubble_sort(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("BubbleSort(%v) = %v; want %v", test.input, test.input, test.expected)
		}
	}
}

func TestBubbleSortConcurrent(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 2, 9, 1, 5, 6}, expected: []int{1, 2, 5, 5, 6, 9}},
		{input: []int{3, 0, -1, 8, 7, 2}, expected: []int{-1, 0, 2, 3, 7, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: generateLargeArray(10000), expected: generateSortedLargeArray(10000)},
	}

	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)

		start := time.Now()
		done := make(chan bool)
		go func() {
			algorithm.Bubble_sort(inputCopy)
			done <- true
		}()

		<-done
		duration := time.Since(start)

		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("BubbleSortConcurrent(%v) = %v; want %v", test.input, inputCopy, test.expected)
		}

		fmt.Printf("BubbleSortConcurrent(%v) took %v\n", test.input, duration)
	}
}

func generateLargeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	return arr
}

func generateSortedLargeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}
	return arr
}
