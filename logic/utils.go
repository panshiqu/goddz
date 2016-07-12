package logic

// Contain 是否包含
func Contain(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

// Index 确认位置
func Index(s []int, i int) int {
	for k, v := range s {
		if v == i {
			return k
		}
	}

	return -1
}

// Count 计算数量
func Count(s []int, i int) int {
	c := 0
	for _, v := range s {
		if v == i {
			c++
		}
	}

	return c
}
