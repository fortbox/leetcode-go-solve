package main

func increasingTriplet(nums []int) bool {
	for i := 1; i < len(nums)-1; i++ {
		arr1 := nums[0:i]
		arr2 := nums[i+1:]
		m := getMin(arr1, nums[i])
		n := getMax(arr2, nums[i])
		if m < nums[i] && nums[i] < n {
			return true
		}
	}
	return false

}

func getMin(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		if k > arr[i] {
			return arr[i]
		}
	}
	return k
}
func getMax(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		if k < arr[i] {
			return arr[i]
		}
	}
	return k
}
