package main

import "fmt"

func main () {
	fmt.Println(printNumWith2(12.3456)) // expect 12.35
	fmt.Println(printBytes([]byte{71, 123}));
}

// print number with two decimals
func printNumWith2(floatIn float64) string {
	return fmt.Sprintf("%.2f", floatIn)
}

// print []byte as hex eg. human readable checksum
func printBytes(data []byte) string {
	return fmt.Sprintf("%x", data)
}