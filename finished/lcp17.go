/*
 * Copyright (c) 2020
 * @Author: xiaoweixiang
 */
package main

import "fmt"

func calculate(s string) int {
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}
func main() {
	var s string = "AB"
	res := calculate(s)
	fmt.Println(res)
}
