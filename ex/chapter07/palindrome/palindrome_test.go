package palindrome

import (
	"sort"
	"testing"
)

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
