package utils

import (
	"math/rand"
	"time"
)

// Queue 先进先出队列
type Queue []string

// Push 向队列中添加一个元素
func (q *Queue) Push(v string) {
	*q = append(*q, v)
}

// Pop 从队列中删除第一个元素
func (q *Queue) Pop() string {
	head := (*q)[0]
	if len(*q) == 1 {
		*q = Queue{}
	} else {
		*q = (*q)[1:]
	}
	return head
}

// IsEmpty 如果队列为空，则返回true
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Stack 先进后出栈
type Stack []string

// Push 向栈中添加一个元素
func (q *Stack) Push(v string) {
	*q = append(*q, v)
}

// Pop 删除栈顶元素
func (q *Stack) Pop() string {
	head := (*q)[len(*q)-1]
	if len(*q) == 1 {
		*q = Stack{}
	} else {
		*q = (*q)[0 : len(*q)-1]
	}

	return head
}

// IsEmpty 如果栈为空，则返回true
func (q *Stack) IsEmpty() bool {
	return len(*q) == 0
}

//GenRandNoRepeated 生成count个[start,end)结束的不重复的随机数
func GenRandNoRepeated(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	vis := map[int]bool{} // 查重
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		if !vis[num] {
			nums = append(nums, num)
			vis[num] = true
		}
	}

	return nums
}
