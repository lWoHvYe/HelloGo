package main

import "fmt"

type ChurchNumeral func(func(interface{}) interface{}) interface{}

func Zero() ChurchNumeral {
	return func(f func(interface{}) interface{}) interface{} {
		return func(x interface{}) interface{} {
			return x
		}
	}
}

func AddOne(n ChurchNumeral) ChurchNumeral {
	return func(f func(interface{}) interface{}) interface{} {
		return func(x interface{}) interface{} {
			return f(n(f).(func(interface{}) interface{})(x))
		}
	}
}

func ToInt(n ChurchNumeral) interface{} {
	return n(func(x interface{}) interface{} {
		return x.(int) + 1
	}).(func(interface{}) interface{})(0)
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
