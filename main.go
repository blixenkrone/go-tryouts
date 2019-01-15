package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 123
	pointerTest(&a)
	x := 1.5
	val := square(&x)
	fmt.Println(val)
}

func square(x *float64) float64 {
	*x = *x * *x
	return *x
}

func pointerTest(args *int) int {
	fmt.Println(*args)
	return *args
}

func mapTest() {
	vals := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range vals {
		fmt.Println(k, v)
	}
}

/* defer function */
func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func getArrLength(a []string) (string, error) {
	if len(a) < 0 {
		return "", errors.New("Not array")
	}

	for idx, val := range a {
		fmt.Println(idx, val)
		if idx == 1 {
			return "Found the one", nil
		}
	}
	return "", nil
}

func getListNumsIdxLen(slice [6]float32) (int, error) {
	if len(slice) < 0 {
		return 0, errors.New("This is not an array")
	}
	fmt.Println(slice[0])
	return 0, nil
}

func addSliceNums(slice []int) (int, error) {
	sum := 0
	if len(slice) < 0 {
		return 0, errors.New("This is not an array")
	}
	for _, val := range slice {
		sum += val
	}
	return sum, nil
}

func substractNums(args ...int) int {
	sum := 0
	for _, val := range args {
		sum -= val
	}
	return sum
}

/**
 *
 *
 *
 */
