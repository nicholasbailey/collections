package collections

import (
	"encoding/binary"
	"hash/maphash"
)

type hashKeyType = uint32

const bitsPerTrieDepth hashKeyType = 5
const sizeOfSlices hashKeyType = 32
const bitMask hashKeyType = sizeOfSlices - 1

type HashMap struct {
	seed maphash.Seed
	size int
	root HAMTNode
}

func NewHashMap() *HashMap {
	seed := maphash.MakeSeed()
	root := &SliceNode{
		size: 0,
		data: make([]HAMTNode, sizeOfSlices),
	}
	return &HashMap{
		seed: seed,
		size: 0,
		root: root,
	}
}

func (hashMap *HashMap) Get(key interface{}) (interface{}, bool) {
	hash := getHash(key, hashMap.seed)
	return hashMap.root.get(hash, 1, key)
}

func (hashMap *HashMap) Set(key interface{}, value interface{}) *HashMap {
	hash := getHash(key, hashMap.seed)
	newRoot, howManyAdded := hashMap.root.set(hash, 1, key, value)
	return &HashMap{
		seed: hashMap.seed,
		size: hashMap.size + howManyAdded,
		root: newRoot,
	}
}

type HAMTNode interface {
	set(hash hashKeyType, depth hashKeyType, key interface{}, value interface{}) (HAMTNode, int)
	get(hash hashKeyType, depth hashKeyType, key interface{}) (interface{}, bool)
}

type SliceNode struct {
	size int
	data []HAMTNode
}

type KeyValueNode struct {
	originalHash hashKeyType
	key          interface{}
	value        interface{}
}

func (node *KeyValueNode) get(hash hashKeyType, depth hashKeyType, key interface{}) (interface{}, bool) {
	if hash != node.originalHash {
		return nil, false
	}
	if key != node.key {
		return nil, false
	}
	return node.value, true
}

func (node *KeyValueNode) set(hash hashKeyType, depth hashKeyType, key interface{}, value interface{}) (HAMTNode, int) {
	// Handle a Key overwrite.
	if node.originalHash == hash {
		if node.key == key {
			if node.value == value {
				return node, 0
			}

			return &KeyValueNode{
				originalHash: hash,
				key:          key,
				value:        value,
			}, 0
		} else {
			panic("Need to support a hash collision node")
		}
	}

	data := make([]HAMTNode, sizeOfSlices)

	selfIndex := getIndexForHash(node.originalHash, depth)
	newIndex := getIndexForHash(hash, depth)
	if selfIndex != newIndex {
		data[selfIndex] = node
		data[newIndex] = &KeyValueNode{
			originalHash: hash,
			key:          key,
			value:        value,
		}
		return &SliceNode{
			data: data,
			size: 2,
		}, 1
	}
	newNode, howManyAdded := node.set(hash, depth+1, key, value)
	data[selfIndex] = newNode
	return &SliceNode{
		data: data,
		size: 1,
	}, howManyAdded
}

func (node *SliceNode) get(hash hashKeyType, depth hashKeyType, key interface{}) (interface{}, bool) {
	index := getIndexForHash(hash, depth)
	target := node.data[index]
	if target == nil {
		return nil, false
	}
	return target.get(hash, depth+1, key)
}

func (node *SliceNode) set(hash hashKeyType, depth hashKeyType, key interface{}, value interface{}) (HAMTNode, int) {
	index := getIndexForHash(hash, depth)
	target := node.data[index]
	if target == nil {
		return &SliceNode{
			size: node.size + 1,
			data: cloneAndSet(node.data, index, &KeyValueNode{
				originalHash: hash,
				key:          key,
				value:        value,
			}),
		}, 1
	}
	newNode, howManyAdded := target.set(hash, depth+1, key, value)
	return &SliceNode{
		size: node.size,
		data: cloneAndSet(node.data, index, newNode),
	}, howManyAdded
}

func getIndexForHash(hash hashKeyType, depth hashKeyType) hashKeyType {
	return hash >> depth * bitsPerTrieDepth & bitMask
}

func getHash(v interface{}, seed maphash.Seed) hashKeyType {
	var h maphash.Hash
	h.SetSeed(seed)
	switch v := v.(type) {
	case string:
		h.WriteString(v)
	case int32:
		buffer := make([]byte, 4)
		binary.LittleEndian.PutUint32(buffer, uint32(v))
		h.Write(buffer)
	case uint32:
		buffer := make([]byte, 4)
		binary.LittleEndian.PutUint32(buffer, v)
		h.Write(buffer)
	case int64:
		buffer := make([]byte, 8)
		binary.LittleEndian.PutUint64(buffer, uint64(v))
		h.Write(buffer)
	case uint64:
		buffer := make([]byte, 8)
		binary.LittleEndian.PutUint64(buffer, v)
		h.Write(buffer)
	default:
		panic(ErrUnhashableType)
	}

	return hashKeyType(h.Sum64())
}

func cloneAndSet(data []HAMTNode, index hashKeyType, node HAMTNode) []HAMTNode {
	newSlice := make([]HAMTNode, sizeOfSlices)
	copy(newSlice, data)
	newSlice[index] = node
	return newSlice
}
