package letter_combinations_of_a_phone_number

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 * Letter Combinations of a Phone Number
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (47.16%)
 * Likes:    457
 * Dislikes: 0
 * Total Accepted:    48.2K
 * Total Submissions: 94.6K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {

	//定义数组
	arr := [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	result := make([]string, 0)

	dfs(arr, 0, digits, &result, "")
	return result
}

// @lc code=end

func dfs(strMap [][]string, start int, digits string, result *[]string, ans string) {
	if start == len(digits) {
		*result = append(*result, ans)
		return
	}

	str := strMap[int(digits[start]-'0')]
	for i := range str {
		temp := ans
		ans += str[i]
		dfs(strMap, start+1, digits, result, ans)
		ans = temp
	}
}
