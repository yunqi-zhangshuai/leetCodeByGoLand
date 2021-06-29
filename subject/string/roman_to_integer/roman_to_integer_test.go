package roman_to_integer

import (
	"leetCodeByGoLand/util/testTool"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	testTool.EqualTest(t, RomanToInt("III"), 3)
	testTool.EqualTest(t, RomanToInt("IX"), 9)
	testTool.EqualTest(t, RomanToInt("MCMXCIV"), 1994)
	testTool.EqualTest(t, RomanToInt("IV"), 4)
	testTool.EqualTest(t, RomanToInt("VX"), 5)
	testTool.EqualTest(t, RomanToInt("XV"), 15)
}
