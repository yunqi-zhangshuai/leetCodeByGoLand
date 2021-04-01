package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSumMap(t *testing.T) {
	numsArr := []int{1, 2, 3, 7, 9, 10, 5}
	// map
	fmt.Println(twoSumMap(numsArr, 9)) //O(n)
	// 暴力枚举
	fmt.Println(twoSum(numsArr, 9)) //O(n²)
}
