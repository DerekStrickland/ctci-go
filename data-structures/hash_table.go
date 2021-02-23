package datastructures

import (
	"fmt"
	"strings"
)

// NaiveHash can be used to inject a function with high probability of key collision
// so that tests can ensure HashTables that are supposed to handle collision do so.
func NaiveHash(value string) int {
	sum := 0

	for _, chr := range value {
		sum = sum + int(chr)
	}

	// fmt.Println(fmt.Sprintf("NaiveHash calculated sum %d for %s", sum, value))
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
	members map[int]*LinkedString
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
		members: make(map[int]*LinkedString),
		hashFunc: hashFn,
	}
}

func (h OpenHashTable) String() string {
	builder := strings.Builder{}

	builder.WriteString("{\n\tmembers: {\n")

	for k, m := range h.members {
		builder.WriteString(fmt.Sprintf("\t\t{ key: %d, member: %+v }\n", k, m))
		n := m.Next

		for n != nil {
			builder.WriteString(fmt.Sprintf("\t\t{ key: %d, member: %+v }\n", k, n))
			n = n.Next
		}
	}

	builder.WriteString("\t}\n}")
	return builder.String()
}

func (h *OpenHashTable) Hash(value string) int {
	return h.hashFunc(value)
}

func (h *OpenHashTable) Add(value string) {
	key := h.hashFunc(value)

	// if no member for this key insert first node
	if el, ok := h.members[key]; !ok {
		h.members[key] = &LinkedString{
			Prev: nil,
			Next: nil,
			Value: value,
		}
	} else { // walk list and link to last node
		prev := el
		next := prev.Next

		//fmt.Println(fmt.Sprintf("prev: %+v", prev))
		//fmt.Println(fmt.Sprintf("next: %+v", next))

		for next != nil {
			prev = next
			next = next.Next
			//fmt.Println(fmt.Sprintf("for prev: %+v", prev))
			//fmt.Println(fmt.Sprintf("for next: %+v", next))
		}

		next = &LinkedString{
			Value: value,
			Prev: prev,
			Next: nil,
		}
		fmt.Println(fmt.Sprintf("final next: %+v", next))

		prev.Next = next
		fmt.Println(fmt.Sprintf("final prev: %+v", prev))
	}
}

func (h *OpenHashTable) Find(value string) *string {
	key := h.hashFunc(value)

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

type ClosedHashTable struct {
	capacity int
	members map[int]string
	hashFunc func(string) int
}

func NewClosedHashTable(capacity int, f func(string) int) ClosedHashTable {
	return ClosedHashTable{
		capacity: capacity,
		members: make(map[int]string, capacity),
		hashFunc: f,
	}
}

func (h *ClosedHashTable) Add(value string) {
	key := h.hashFunc(value)
	_, found := h.members[key]

	for found {
		key = (key + 1) % h.capacity
		_, found = h.members[key]
	}

	h.members[key] = value
}

func (h *ClosedHashTable) Find(value string) *string {
	key := h.hashFunc(value)
	el, found := h.members[key]

	for found {
		if el == value {
			return &el
		}

		key = (key + 1) % h.capacity
		el, found = h.members[key]
	}

	return nil
}

type QuadraticHashTable struct {
	capacity int
	members map[int]string
	hashFunc func(string)int
}

func NewQuadraticHashTable(capacity int, hashFn func(string)int) QuadraticHashTable {
	return QuadraticHashTable{
		capacity: capacity,
		members: make(map[int]string, capacity),
		hashFunc: hashFn,
	}
}

func (h *QuadraticHashTable) quadFunc(key, factor int) int {
	return (key + (factor * factor)) % h.capacity
}

func (h *QuadraticHashTable) Add(value string) {
	key := h.hashFunc(value)
	factor := 0

	_, found := h.members[key]
	for found {
		factor++
		key = h.quadFunc(key, factor)
		_, found = h.members[key]
	}

	h.members[key] = value
}

func (h *QuadraticHashTable) Find(value string) *string {
	key := h.hashFunc(value)
	factor := 0

	el, found := h.members[key]

	for found {
		if el == value {
			return &el
		}
		factor++
		key = h.quadFunc(key, factor)
		el, found = h.members[key]
	}

	return nil
}