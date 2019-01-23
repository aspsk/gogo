package gramana

func gramana(s1, s2 string) bool {

	m1 := make(map[rune]int, 0)
	for _, r := range s1 {
		m1[r]++
	}

	m2 := make(map[rune]int, 0)
	for _, r := range s2 {
		m2[r]++
	}

	for key, value := range m1 {
		if m2[key] != value {
			return false
		}
	}

	return true
}
