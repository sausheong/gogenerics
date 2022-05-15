package generics

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func Add[T int | float64](a, b T) T {
	return a + b
}

func AddNumbers[T Number](a, b T) T {
	return a + b
}
