package main

import "fmt"

type Number interface {
	// ~ means addition of subtypes to specified type
	~int | int32 | int64 | float32 | float64
}

func calcSum[T Number](a, b T) T {
	var res T
	res = a + b
	return res
}

func genericWithTypeCheck[T Number](a, b T) T {
	switch any(a).(type) {
	case int:
		return a + b
	case float64:
		return a * b
	default:
		return a - b
	}
}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(a T) {
	s.data = append(s.data, a)
}

func (s *Stack[T]) Pop() T {
	n := len(s.data) - 1
	if n < 0 {
		var res T
		return res
	}
	res := s.data[n]
	s.data = s.data[:n]
	return res
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	set := make(map[T]struct{}, len(elems))
	for _, elem := range elems {
		set[elem] = struct{}{}
	}
	return set
}

func (s Set[T]) Put(elem T) {
	s[elem] = struct{}{}
}
func (s Set[T]) Has(elem T) bool {
	_, has := s[elem]
	return has
}

func main() {
	fmt.Println("Testing generic functions")
	fmt.Println(calcSum(1, 2))
	fmt.Println(calcSum[int](1, 2)) // Use explicit type
	fmt.Println(calcSum(1.5, 2.5))

	fmt.Println("Testing generic functions with type checking")
	fmt.Println(genericWithTypeCheck(2, 4))
	fmt.Println(genericWithTypeCheck[int](2, 4)) // Use explicit type
	fmt.Println(genericWithTypeCheck(1.5, 2.5))

	fmt.Println("Testing generic types (Stack of any values")
	stackAny := Stack[any]{}
	stackAny.Push(1)
	stackAny.Push(1.2)
	stackAny.Push("asd")
	fmt.Println(stackAny.Pop())
	fmt.Println(stackAny.Pop())
	fmt.Println(stackAny.Pop())
	fmt.Println(stackAny.Pop())

	fmt.Println("Testing Stack of int")
	stackInt := Stack[int]{}
	stackInt.Push(1)
	stackInt.Push(2)
	//stackInt.Push(1.2) // Error
	fmt.Println(stackInt.Pop())
	fmt.Println(stackInt.Pop())
	fmt.Println(stackInt.Pop())

	fmt.Println("Testing Set type")
	set := NewSet("A", "B")
	set.Put("C")
	fmt.Println(set)
	fmt.Println(set.Has("Z"))
	fmt.Println(set.Has("C"))
}
