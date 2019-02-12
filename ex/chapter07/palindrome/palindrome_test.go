package palindrome

import (
	"sort"
	"time"
	"math/rand"
	"testing"
)

func TestPalindromeString(t *testing.T) {

	s1 := "абба"
	s2 := "лезунасанузел"
	s3 := "ямайка"
	s4 := "палиндром"

	if !IsStringPalindrome(s1) {
		t.Error("s1 is a palindrome")
	}

	if !IsStringPalindrome(s2) {
		t.Error("s1 is a palindrome")
	}

	if IsStringPalindrome(s3) {
		t.Error("s3 is not a palindrome")
	}

	if IsStringPalindrome(s4) {
		t.Error("s4 is not a palindrome")
	}

}

func TestPalindrome(t *testing.T) {

	s1 := sort.StringSlice([]string{"123", "456", "456", "123"})
	s2 := sort.StringSlice([]string{"123", "456", "789", "456", "123"})
	s3 := sort.StringSlice([]string{"123", "456", "789", "abc", "123"})
	s4 := sort.StringSlice([]string{"123", "456", "abc", "123"})

	if !IsPalindrome(s1) {
		t.Error("s1 is a palindrome")
	}

	if !IsPalindrome(s2) {
		t.Error("s1 is a palindrome")
	}

	if IsPalindrome(s3) {
		t.Error("s3 is not a palindrome")
	}

	if IsPalindrome(s4) {
		t.Error("s4 is not a palindrome")
	}

}

// exercise 11.3

func randomNonPalindrome(rng *rand.Rand) string {
	n := 10 + rng.Intn(25) // make it non-empty
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
	}
	return string(runes)
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsStringPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}

// exercise 11.4


func randomPalindromePunct(rng *rand.Rand) string {
	n := rng.Intn(25) + 1 // make it non-empty
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}

	runes2 := make([]rune, n * 2)
	for i := 0; i < n; i++ {
		runes2[2*i] = runes[i]
		runes2[2*i+1] = rune(' ')
	}

	return string(runes2)
}

func TestRandomPalindromesPunct(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindromePunct(rng)
		if !IsPalindromeBook(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

// below is the code taken from the book, so

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsStringPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
