package linked_list_cycle

import (
	"leetCodeByGoLand/dataStructure/linkedList/linkedList"
)

//给定一个链表，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//如果链表中存在环，则返回 true 。 否则，返回 false 。
//
//
//
//进阶：
//
//你能用 O(1)（即，常量）内存解决此问题吗？
//
//
//
//示例 1：
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
//示例 2：
//
//
//
//输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
//示例 3：
//
//
//
//输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
//
//
//提示：
//
//链表中节点的数目范围是 [0, 104]
//-105 <= Node.val <= 105
//pos 为 -1 或者链表中的一个 有效索引 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/linked-list-cycle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// hasCycleFastSlowPoint
// 快慢指针
// 1. 初始化 fast ， slow 两个指针
// 2. 跳过头节点判断
// 3. fast slow 不能为 nil
// 4. fast 走两步，slow 走一步
// 5. fast == slow 有环
func hasCycleFastSlowPoint(head *linkedList.SingleNode) bool {
	fast, slow := head, head
	var index uint32
	for fast != nil && slow != nil {
		index++
		if fast == slow && index > 1 {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

// hasCycleByMap
//  集合判重
// 1. 定义一个key 是节点指针地址的map
// 2. 遍历链表,判断节点指针是否存在于map中
// 3. 节点指针不存在,将node 指针放置map中;存在返回
func hasCycleByMap(head *linkedList.SingleNode) bool {

	pointMap := make(map[*linkedList.SingleNode]uint32)
	cur := head
	var index uint32
	for cur != nil {
		index++
		if _, ok := pointMap[cur]; ok {
			return true
		}
		pointMap[cur] = index
		cur = cur.Next
	}
	return false
}
