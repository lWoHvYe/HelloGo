package main

import "fmt"

type FuncL1 func(any) any
type FuncL2 func(FuncL1) FuncL1

type ChurchNumeral FuncL2

func Zero() ChurchNumeral {
	return func(f FuncL1) FuncL1 {
		return func(x any) any {
			return x
		}
	}
}

func AddOne(n ChurchNumeral) ChurchNumeral {
	return func(f FuncL1) FuncL1 {
		return func(x any) any {
			return f(n(f)(x))
		}
	}
}

func ToInt(n ChurchNumeral) any {
	return n(func(x any) any {
		num, ok := x.(int)
		if ok {
			return num + 1
		}
		panic("Cur input is not a num")
	})(0)
}

func main() {
	zero := Zero()
	one := AddOne(zero)
	two := AddOne(one)
	three := AddOne(two)

	fmt.Println("Zero:", ToInt(zero))
	fmt.Println("One:", ToInt(one))
	fmt.Println("Two:", ToInt(two))
	fmt.Println("Three:", ToInt(three))

}
