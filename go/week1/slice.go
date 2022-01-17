package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}

	// expect 5, 1, 2, 4, 7
	s = Add(s, 0, 5) 
	fmt.Println(s);

	// expect 5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Println(s);

	// expect 5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Println(s);

	// expect 5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Println(s);

	// expect 9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Println(s);

	// expect 9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Println(s);

}

func Add(s []int, index int, value int) []int {
	
	if index < 0 || index > len(s) {
		// TODO: throw
	}

	// gotcha: re-slicing a slice doesn’t make a copy of the underlying array
	// tmp := append(s[0:index], value) // shared slice/reference - s => [5,2,4,7]
	// fmt.Println(tmp);

	tmp := make([]int, 0, len(s) + 1)
	tmp = append(tmp, s[:index]...)
	tmp = append(tmp, value)
	tmp = append(tmp, s[index:]...)
	return tmp
}

func Delete(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		// TODO: throw
	}

	// seems fine with s[:len(s)] <--- there is an undefined/end element?
	// s[len(s)] => []
	
	// if index == len(s) - 1 {
	// 	return s[:index]
	// }

	return append(s[:index], s[index+1:]...);;
}