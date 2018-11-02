package main

import "fmt"

/**
   两数之和

  给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
  你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
  示例:

 给定 nums = [2, 7, 11, 15], target = 9
 因为 nums[0] + nums[1] = 2 + 7 = 9
 所以返回 [0, 1]
 */

func main() {

	nums_arr := []int{1, 2, 3, 7, 9, 10, 5}
	fmt.Println(_twoSum(nums_arr, 9)) //O(n)
	fmt.Println(twoSum(nums_arr, 9))  //O(n²)
}

/**
思路1暴力枚举
双层循环,
第一层循环计算 给定数与数组当前循环值的差
第二层循环从,数组下一个值判断是不是和第一层循环的差值相等,并返回两个数index

O(n²)
*/

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		wz := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == wz {
				return []int{i, j}
			}
		}

	}

	return []int{};
}

/*
思路2
在循环中计算给定数与数组中数的差,

将value为index,index为value放置到map中

判断当前差值有没有在map中
0(n)
*/
func _twoSum(nums []int, target int) []int {

	tmpMap := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if n, ok := tmpMap[complement]; ok {
			return []int{n, i}
		}
		tmpMap[v] = i
	}

	return []int{}

}
