package main

import "fmt"

type UnaryFunc func(any) any
type HighOrderFunc func(UnaryFunc) UnaryFunc
type ChurchNumber HighOrderFunc

func Zero() ChurchNumber {
	return func(f UnaryFunc) UnaryFunc {
		return func(x any) any {
			return x
		}
	}
}

func Increment(n ChurchNumber) ChurchNumber {
	return func(f UnaryFunc) UnaryFunc {
		return func(x any) any {
			return f(n(f)(x))
		}
	}
}

func ToInteger(n ChurchNumber) any {
	return n(func(x any) any {
		if num, ok := x.(int); ok {
			return num + 1
		}
		panic("Not a number")
	})(0)
}

func DoubleStr(n ChurchNumber) any {
	return n(func(x any) any {
		if str, ok := x.(string); ok {
			return str + "+" + str
		}
		panic("Not a string")
	})("-")
}

func main() {
	churchZero := Zero()
	churchOne := Increment(churchZero)
	churchTwo := Increment(churchOne)
	churchThree := Increment(churchTwo)

	fmt.Println("Zero:", ToInteger(churchZero))
	fmt.Println("One:", ToInteger(churchOne))
	fmt.Println("Two:", ToInteger(churchTwo))
	fmt.Println("Three:", ToInteger(churchThree))

	fmt.Println("Zero:", DoubleStr(churchZero))
	fmt.Println("One:", DoubleStr(churchOne))
	fmt.Println("Two:", DoubleStr(churchTwo))
	fmt.Println("Three:", DoubleStr(churchThree))
}
