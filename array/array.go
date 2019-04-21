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

// 904
func totalFruit(tree []int) int {

	if len(tree) < 3 {
		return len(tree)
	}
	a := tree[0]
	global, local, l := 0, 0, 1
	var j, b int
	for t := 1; t < len(tree); t++ {
		if tree[t] != a {
			j = t
			b = tree[t]
			global, local = t + 1, t + 1
			break
		}
	}
	if j == 0 {
		return len(tree)
	}
	for k := j + 1; k < len(tree); k++ {
		// a b b
		if tree[k] == b {
			local++
			l++
		} else if tree[k] == a { // a b a
			local++
			l = 1
			a, b = b, a
		} else { // a b c
			a, b = b, tree[k]
			local, l = l + 1, 1
		}
		if local > global {
			global = local
		}
	}


	return global
}
