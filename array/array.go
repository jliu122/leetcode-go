package array

// 001
func twoSum(nums []int, target int) []int {

	result := make([]int, 2)
	if len(nums) < 2 {
		return result
	}

	indexMap := make(map[int]int)
	for i, n := range nums {
		if v, ok := indexMap[target - n]; ok {
			result[0] = v
			result[1] = i
			return result
		}
		indexMap[n] = i
	}
	return result
}
