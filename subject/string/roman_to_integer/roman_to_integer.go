package roman_to_integer

// RomanToInt
// 罗马数字转整数
func RomanToInt(s string) int {

	selectValue := func(char string) int {
		var val int
		switch char {
		case "I":
			val = 1
		case "V":
			val = 5
		case "X":
			val = 10
		case "L":
			val = 50
		case "C":
			val = 100
		case "D":
			val = 500
		case "M":
			val = 1000
		}

		return val
	}

	sum, preval := 0, 0
	for _, luoma := range s {
		num := selectValue(string(luoma))
		if preval < num {
			sum -= preval
		} else {
			sum += preval
		}
		preval = num
	}
	sum += preval
	return sum
}
