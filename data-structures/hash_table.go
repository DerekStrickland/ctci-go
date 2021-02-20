package datastructures

import (
	"fmt"
)

// NaiveHash can be used to inject a function with high probability of key collision
// so that tests can ensure HashTables that are supposed to handle collision do so.
func NaiveHash(value string) int {
	sum := 0

	for _, chr := range value {
		sum = sum + int(chr)
	}

	return sum
}

// reusable has function across HashTable implementations.
func hash(value string) int {
	// fmt.Println(fmt.Sprintf("evaluating hash for string %s", value))
	sum := 0

	for i, chr := range value {
		chrInt := int(chr)
		// fmt.Println(fmt.Sprintf("chrInt is %d", chrInt))
		sum = sum + (chrInt * (i + 1))
		// fmt.Println(fmt.Sprintf("sum is %d", sum))
	}

	hash := sum % 2069
	// fmt.Println(fmt.Sprintf("hash for value %s evaluated to %d", value, hash))
	return hash
}

// HashTable is a naive HashTable with no collision detection
type HashTable struct {
	members map[int]string
}

// NewHashTable is a factory method for the HashTable type
func NewHashTable() HashTable {
	return HashTable{
		members: make(map[int]string),
	}
}

func (h *HashTable) String() string {
	return fmt.Sprintf("%+v", h.members)
}

func (h *HashTable) Hash(value string) int {
	return hash(value)
}

func (h *HashTable) Add(value string) {
	h.members[h.Hash(value)] = value
}

func (h *HashTable) Exists(value string) bool {
	return h.Find(value) != ""
}

func (h * HashTable) Find(value string) string {
	return h.members[h.Hash(value)]
}

// OpenHashTable uses separate chaining (open hashing) to manage key collisions
// as a set of linked lists
type OpenHashTable struct {
	members map[int]LinkedString
	hashFunc func(string) int
}

type LinkedString struct {
	Prev *LinkedString
	Next *LinkedString
	Value string
}

func NewOpenHashTable(hashFn func(string)int) OpenHashTable {
	if hashFn == nil {
		hashFn = hash
	}

	return OpenHashTable{
		members: make(map[int]LinkedString),
		hashFunc: hashFn,
	}
}

func (h *OpenHashTable) String() string {
	return fmt.Sprintf("%+v", h.members)
}

func (h *OpenHashTable) Hash(value string) int {
	return hash(value)
}

func (h *OpenHashTable) Add(value string) {
	key := hash(value)

	// if no member for this key insert first node
	if el, ok := h.members[key]; !ok {
		h.members[key] = LinkedString{
			Prev: nil,
			Next: nil,
			Value: value,
		}
	} else { // walk list and link to last node
		next := el.Next
		prev := &el

		for next != nil {
			prev = next
			next = next.Next
		}

		prev.Next = &LinkedString{
			Value: value,
			Prev: prev,
			Next: nil,
		}
	}
}

func (h *OpenHashTable) Find(value string) *string {
	key := h.Hash(value)

	if el, ok := h.members[key]; ok {
		if el.Value == value {
			return &el.Value
		}

		next := el.Next

		for next != nil {
			if next.Value == value {
				return &next.Value
			}
			next = next.Next
		}
	}

	return nil
}
