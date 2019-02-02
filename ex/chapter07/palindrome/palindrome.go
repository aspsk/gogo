package palindrome

import "sort"

func IsPalindrome(s sort.Interface) bool {
	for n, i := s.Len(), 0; i < n/2; i++ {
		if s.Less(i, n-i-1) || s.Less(n-i-1, i) {
			return false
		}
	}
	return true
}
