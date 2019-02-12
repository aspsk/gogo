package palindrome

import (
	"sort"
	"unicode"
)

func IsPalindrome(s sort.Interface) bool {
	for n, i := s.Len(), 0; i < n/2; i++ {
		if s.Less(i, n-i-1) || s.Less(n-i-1, i) {
			return false
		}
	}
	return true
}

func IsStringPalindrome(s string) bool {
	r := []rune(s)

	for n, i := len(r), 0; i < n/2; i++ {
		if r[i] != r[n-i-1] {
			return false
		}
	}

	return true
}

func IsPalindromeBook(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
