package main

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.nums
}
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	buf := make([]int, len(this.nums))
	copy(buf, this.nums)
	for i := range nums {
		j := rand.Intn(len(buf))
		nums[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return nums
}
