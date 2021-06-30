package add_two_numbers

import (
	"fmt"
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	test1_l1 := linkedList.NewSingleList(8, 2, 3, 7)
	test1_l2 := linkedList.NewSingleList(3, 3, 2, 0)
	// 7 3 2  8
	// 0 2 3  3
	// ==
	// 7 5 6  1
	// 反转
	// 1 6 5 7
	test1_want := linkedList.NewSingleList(1, 6, 5, 7)

	type args struct {
		l1 *linkedList.SingleNode
		l2 *linkedList.SingleNode
	}
	tests := []struct {
		name string
		args args
		want *linkedList.SingleNode
	}{
		{
			name: "test1",
			args: args{
				l1: test1_l1.GetHead(),
				l2: test1_l2.GetHead(),
			},
			want: test1_want.GetHead(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := addTwoNumbers(tt.args.l1, tt.args.l2)
			fmt.Println(tt.name + "※" + strings.Repeat("-", 50))
			head := node
			for head != nil {
				fmt.Println(head.Item)
				head = head.Next
			}

		})
	}
}
