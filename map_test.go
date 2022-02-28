package collections

import (
	"testing"
)

func TestHashMapSet(t *testing.T) {
	expect := expectFor(t)
	m0 := NewHashMap()
	val, found := m0.Get("hello")
	expect(val).ToBe(nil)
	expect(found).ToBe(false)

	m1 := m0.Set("hello", 5)
	val, found = m1.Get("hello")
	expect(val).ToBe(5)
	expect(found).ToBe(true)

	m2 := m1.Set("hello", 8)

	val, found = m2.Get("hello")
	expect(val).ToBe(8)
	expect(found).ToBe(true)

	val, found = m1.Get("hello")
	expect(val).ToBe(5)
	expect(found).ToBe(true)

	val, found = m0.Get("hello")
	expect(val).ToBe(nil)
	expect(found).ToBe(false)
}

func TestHashMapManyValues(t *testing.T) {
	expect := expectFor(t)
	random := fakerFor(t)
	hashMap := NewHashMap()
	goMap := map[string]string{}
	for i := 0; i < 2000; i++ {
		key := random.String(4, 30)
		value := random.String(4, 30)
		goMap[key] = value
		hashMap = hashMap.Set(key, value)
	}
	for k, v := range goMap {
		val, found := hashMap.Get(k)
		expect(val).ToBe(v)
		expect(found).ToBe(true)
	}
}
