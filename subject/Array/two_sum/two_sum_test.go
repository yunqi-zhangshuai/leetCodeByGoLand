package two_sum

import (
	"fmt"
	"testing"
)

var numsArr = []int{1, 2, 3, 7, 9, 10, 5}

func TestTwoSumMap(t *testing.T) {
	// map
	fmt.Println(twoSumMap(numsArr, 9)) //O(n)

}

func TestTwoSum(t *testing.T) {
	// 暴力枚举
	fmt.Println(twoSum(numsArr, 9)) //O(n²)
}
