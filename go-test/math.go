package main

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

// WrongMulti intends to multiply a and b,
// however, the implementation is wrong and would not pass the unit test
func WrongMulti(a int, b int) int {
	return a / b
}
