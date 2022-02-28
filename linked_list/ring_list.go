package main

type Ring struct {
	next, prev *Ring       // 前驱和后驱节点
	Value      interface{} // 数据
}

// 初始化空的循环链表，前驱和后驱都指向自己，因为时循环的
// 因为绑定的前驱和后驱节点为自己，没有循环，时间复杂度为:O(1)
func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// New 创建一个指定大小N的循环链表，值全为空：
// 会连续绑定前驱和后驱节点，时间复杂度为:O(n)
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 0; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// Next 获取上一个或下一个节点
// 获得前驱或后驱节点，时间复杂度为:O(1)
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// Prev 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

func main() {
	r := new(Ring)
	r.init()
}
