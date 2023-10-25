package lru

func (c *Cache) Get(key string) (Value, bool) {
	if v, ok := c.cache[key]; ok {
		c.ll.MoveToFront(v)
		kv := v.Value.(*entry)

		return kv.value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())

		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())

		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.nBytes += int64(len(key)) + int64(value.Len())
		c.cache[key] = ele
	}

	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int64 {
	return int64(c.ll.Len())
}
