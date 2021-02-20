package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashFuncIsUniqueForSameSet(t *testing.T) {
	require := require.New(t)

	hashTable := NewHashTable()

	values := []string{"abcdefghijklmnop", "ponmlkjihgfedcba"}

	require.True(hashTable.Hash(values[0]) != hashTable.Hash(values[1]))
}

func TestHashFuncFrequency(t *testing.T) {
	testString := "ababcd"
	hashTable := NewHashTable()
	testChars := make(map[int]int)

	for _, c := range testString {
		ch := string(c)
		// fmt.Println(fmt.Sprintf("Processing test string char %s", ch))
		key := hashTable.Hash(ch)
		// fmt.Println(fmt.Sprintf("Evaluated key of %d", key))
		counter := testChars[key]
		// fmt.Println(fmt.Sprintf("Count of keys in map %d for key", counter))
		counter++
		testChars[key] = counter
	}

	for i := int('a'); i < int('z'); i++ {
		ch := string(rune(i))
		key := hashTable.Hash(ch)
		fmt.Println(fmt.Sprintf("Character %s with key %d has count %d", ch, key, testChars[key]))
	}
}

func TestHashTableAddAndFind(t *testing.T) {
	testStrings := []string{"now is the", "time for all", "good people", "to come to", "the aid of their country"}
	hashTable := NewHashTable()

	for _, s := range testStrings {
		hashTable.Add(s)
	}

	fmt.Println(fmt.Sprintf("HashTable is %+v", hashTable))

	for _, s := range testStrings {
		require.Equal(t, s, hashTable.Find(s), "HashTable failed to find test strings")
	}
}

func TestOpenHashTableAddAndFind(t *testing.T) {
	testStrings := []string{"now is the", "time for all", "good people", "to come to", "the aid of their country", "now is the", "time for all", "good people", "to come to", "the aid of their country"}
	hashTable := NewOpenHashTable(NaiveHash)

	for _, s := range testStrings {
		hashTable.Add(s)
	}

	fmt.Println(fmt.Sprintf("OpenHashTable is %+v", hashTable))

	for _, s := range testStrings {
		require.Equal(t, s, *hashTable.Find(s), "OpenHashTable failed to find test strings")
	}
}
