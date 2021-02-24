package datastructures

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashFuncIsUniqueForSameSet(t *testing.T) {

	hashTable := NewHashTable()

	values := []string{"abcdefghijklmnop", "ponmlkjihgfedcba"}

	require.True(t, hashTable.Hash(values[0]) != hashTable.Hash(values[1]))
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
		count := testChars[key]
		if strings.Index(testString, ch) > - 1 {
			require.Greater(t, count, 0)
		} else {
			require.Equal(t, count, 0)
		}
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
	testStrings := []string{ "dabc", "abcd", "bcda", "cdab", "ghef", "efgh", "fghe", "hefg"}

	hashTable := NewOpenHashTable(NaiveHash)

	for _, s := range testStrings {
		hashTable.Add(s)
	}

	fmt.Println(fmt.Sprintf("OpenHashTable is \n%+v", hashTable))

	for _, s := range testStrings {
		// fmt.Println(fmt.Sprintf("resolving entry for string %s", s))
		require.Equal(t, s, *hashTable.Find(s), "OpenHashTable failed to find test strings")
	}
}

func TestOpenHashNotFound(t *testing.T) {
	hashTable := NewOpenHashTable(NaiveHash)

	result := hashTable.Find("foo")

	require.Nil(t, result, "OpenHashTable failed to evaluate non-member string to nil")
}

func TestClosedHashTableAddAndFind(t *testing.T) {
	testStrings := []string{ "dabc", "abcd", "bcda", "cdab", "ghef", "efgh", "fghe", "hefg"}

	hashTable := NewClosedHashTable(8, NaiveHash)

	for _, s := range testStrings {
		hashTable.Add(s)
	}

	fmt.Println(fmt.Sprintf("ClosedHashTable is \n%+v", hashTable))

	for _, s := range testStrings {
		// fmt.Println(fmt.Sprintf("resolving entry for string %s", s))
		require.Equal(t, s, *hashTable.Find(s), "ClosedHashTable failed to find test strings")
	}
}

func TestClosedHashNotFound(t *testing.T) {
	hashTable := NewClosedHashTable(1, NaiveHash)
	hashTable.Add("oof")

	result := hashTable.Find("foo")

	require.Nil(t, result, "ClosedHashTable failed to evaluate non-member string to nil")
}

func TestQuadraticHashTableAddAndFind(t *testing.T) {
	testStrings := []string{ "dabc", "abcd", "bcda", "cdab", "ghef", "efgh", "fghe", "hefg"}

	hashTable := NewQuadraticHashTable(8, NaiveHash)

	for _, s := range testStrings {
		hashTable.Add(s)
	}

fmt.Println(fmt.Sprintf("QuadraticHashTable is \n%+v", hashTable))

	for _, s := range testStrings {
		// fmt.Println(fmt.Sprintf("resolving entry for string %s", s))
		require.Equal(t, s, *hashTable.Find(s), "QuadraticHashTable failed to find test strings")
	}
}

func TestQuadraticHashNotFound(t *testing.T) {
	hashTable := NewQuadraticHashTable(1, NaiveHash)
	hashTable.Add("oof")

	result := hashTable.Find("foo")

	require.Nil(t, result, "QuadraticHashTable failed to evaluate non-member string to nil")
}

func TestDoubleHashTableAddAndFind(t *testing.T) {
	testStrings := []string{ "dabc", "abcd", "bcda", "cdab", "ghef", "efgh", "fghe", "hefg"}

	hashTable := NewDoubleHashTable(8, NaiveHash)

	for _, s := range testStrings {
		hashTable.Add(s)
	}

	fmt.Println(fmt.Sprintf("DoubleHashTable is \n%+v", hashTable))

	for _, s := range testStrings {
		fmt.Println(fmt.Sprintf("resolving entry for string %s", s))
		require.NotNil(t, hashTable.Find(s), "DoubleHashTable failed to find test strings")
	}
}

func TestDoubleHashNotFound(t *testing.T) {
	hashTable := NewDoubleHashTable(1, NaiveHash)
	hashTable.Add("oof")

	result := hashTable.Find("foo")

	require.Nil(t, result, "DoubleHashTable failed to evaluate non-member string to nil")
}