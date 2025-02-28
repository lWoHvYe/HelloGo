package main

import "strconv"

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
// pointer method 以pointer为接收的方法，如何定义generics
// Setter2 is a type constraint that requires that the type
// implement a Set method that sets the value from a string,
// and also requires that the type be a pointer to its type parameter.
type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

// FromStrings2 takes a slice of strings and returns a slice of T,
// calling the Set method to set each returned value.
//
// We use two different type parameters so that we can return
// a slice of type T but call methods on *T aka PT.
// The Setter2 constraint ensures that PT is a pointer to T.
func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}

// Settable is an integer type that can be set from a string.
type Settable int

// Set sets the value of *p from a string.
func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

func F3() {
	// Here we just pass one type argument.
	nums := FromStrings2[Settable]([]string{"1", "2"})
	// Now nums is []Settable{1, 2}.
	//...
	println(nums)
}

// ----------Using types that refer to themselves in constraints，直接定义泛型需要有哪些方法
// Index returns the index of e in s, or -1 if not found.
func Index[T interface{ Equal(T) bool }](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}

// equalInt is a version of int that implements Equaler.
type equalInt int

// The Equal method lets equalInt implement the Equaler constraint.
func (a equalInt) Equal(b equalInt) bool { return a == b }

// indexEqualInts returns the index of e in s, or -1 if not found.
func indexEqualInt(s []equalInt, e equalInt) int {
	// The type argument equalInt is shown here for clarity.
	// Function argument type inference would permit omitting it.
	return Index[equalInt](s, e)
}

// Omissions
//We believe that this design covers the basic requirements for generic programming. However, there are a number of programming constructs that are not supported.
//
//No specialization. There is no way to write multiple versions of a generic function that are designed to work with specific type arguments.
//No metaprogramming. There is no way to write code that is executed at compile time to generate code to be executed at run time.
//No higher level abstraction. There is no way to use a function with type arguments other than to call it or instantiate it. There is no way to use a generic type other than to instantiate it.
//No general type description. In order to use operators in a generic function, constraints list specific types, rather than describing the characteristics that a type must have. This is easy to understand but may be limiting at times.
//No covariance or contravariance of function parameters.
//No operator methods. You can write a generic container that is compile-time type-safe, but you can only access it with ordinary methods, not with syntax like c[k].
//No currying. There is no way to partially instantiate a generic function or type, other than by using a helper function or a wrapper type. All type arguments must be either explicitly passed or inferred at instantiation time.
//No variadic type parameters. There is no support for variadic type parameters, which would permit writing a single generic function that takes different numbers of both type parameters and regular parameters.
//No adaptors. There is no way for a constraint to define adaptors that could be used to support type arguments that do not already implement the constraint, such as, for example, defining an == operator in terms of an Equal method, or vice-versa.
//No parameterization on non-type values such as constants. This arises most obviously for arrays, where it might sometimes be convenient to write type Matrix[n int] [n][n]float64. It might also sometimes be useful to specify significant values for a container type, such as a default value for elements.
