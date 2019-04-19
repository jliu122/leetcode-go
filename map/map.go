package _map

// 146
type LRUCache struct {

	Cap int
	Len int
	Head *Cache
	Tail *Cache
	Map map[int]*Cache
}

type Cache struct {

	Prev *Cache
	Next *Cache
	Key int
	Val int
}


func Constructor(capacity int) LRUCache {

	m := make(map[int]*Cache)
	head := Cache{}
	tail := Cache{}
	head.Next = &tail
	tail.Prev = &head
	return LRUCache{capacity, 0, &head, &tail, m}
}


func (this *LRUCache) Get(key int) int {

	if c, ok := this.Map[key]; ok {

		this.Remove(c)
		this.SetHead(c)
		return c.Val
	}
	return -1
}

func (this *LRUCache) Remove(c *Cache) {

	prev := c.Prev
	next := c.Next
	prev.Next = next
	next.Prev = prev

}

func (this *LRUCache) SetHead(c *Cache) {

	c.Prev = this.Head
	c.Next = this.Head.Next
	this.Head.Next.Prev = c
	this.Head.Next = c
}


func (this *LRUCache) Put(key int, value int)  {

	if c, ok := this.Map[key]; ok {
		c.Val = value
		this.Remove(c)
		this.SetHead(c)
	} else {
		nc := Cache{Key: key, Val: value}
		this.SetHead(&nc)
		this.Map[key] = &nc
		if this.Len == this.Cap {
			tc := this.Tail.Prev
			delete(this.Map, tc.Key)
			this.Remove(tc)
		} else {
			this.Len++
		}
	}
}
