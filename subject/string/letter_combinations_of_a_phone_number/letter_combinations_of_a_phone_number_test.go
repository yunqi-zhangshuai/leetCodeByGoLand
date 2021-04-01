package letter_combinations_of_a_phone_number

import (
	"fmt"
	"testing"
)

func TestPhoneNumber(t *testing.T) {
	combinations := letterCombinations("4257")
	fmt.Println(combinations)
}
