package main

import (
	"errors"
	"fmt"
)

// Person blabla
type Person struct {
	name     string
	age      byte
	lastName string
	gender   string
	// sayHello ()
}

func main() {
	simon := Person{"simon", 26, "blix", "m"}
	simone := Person{"simone", 27, "schultz", "f"}

	fmt.Println(simon)
	fmt.Println(simone.getMarried("blix"))
	fmt.Println(simone)
}

func (p Person) returnPerson() string {
	return p.name + `is`
}

func (p *Person) getMarried(lName string) string {
	if p.gender == "m" {
		return "its a man!"
	}
	p.lastName = lName
	return "changed" + p.name + " last name to " + lName
}
func (p *Person) changePerson() {
	p.age++
}

func sayHelloFunc(s string) string {
	return s
}

func square(x *float64) float64 {
	*x = *x * *x
	return *x
}

func pointerTest(args *int) int {
	fmt.Println(*args)
	return *args
}

func mapTest(vals map[string]int) {
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
