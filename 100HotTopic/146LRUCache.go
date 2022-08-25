package _00HotTopic

import (
	"container/list"
)

// 使用哈希表 + 双向链表
type LRUCache struct {
	// 使用 list 构造 LRUCache
	cache *list.List
	// 使用 map 映射 list 中的 entry 满足 Get 时间复杂度为 O(1)
	cacheMap map[int]*list.Element
	// length 限制 Cache 的容量
	length int
}

// entry 为 list 中每个节点的存储内容
type entry struct {
	key   int
	value int
}

// Constructor 构造一个新的 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{cache: list.New(), cacheMap: make(map[int]*list.Element, capacity), length: capacity}
}

func (this *LRUCache) Get(key int) int {
	// cacheMap 里存在 key 对应的 entry，直接返回 entry 的 value
	if elem, ok := this.cacheMap[key]; ok {
		// 将该条目移动到链表头部
		this.cache.MoveToFront(elem)
		// 返回相应的值
		kv := elem.Value.(*entry)
		return kv.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 通过 map 判断 LRUCache 中是否有该 entry
	if elem, ok := this.cacheMap[key]; ok {
		// 将该条目移动到链表头部
		this.cache.MoveToFront(elem)
		// 修改条目的 value
		kv := elem.Value.(*entry)
		kv.value = value
	} else {
		// 没有该条目，往 list 插入新 entry
		elem := this.cache.PushFront(&entry{key: key, value: value})
		// 增加该条目的映射
		this.cacheMap[key] = elem
	}
	// Put 后判断是否超过容量
	if this.cache.Len() > this.length {
		// 如果超过容量则进行淘汰
		this.deleteOldest()
	}
}

// deleteOldest 实现缓存淘汰
func (this *LRUCache) deleteOldest() {
	elem := this.cache.Back()
	if elem != nil {
		this.cache.Remove(elem)
		kv := elem.Value.(*entry)
		delete(this.cacheMap, kv.key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
