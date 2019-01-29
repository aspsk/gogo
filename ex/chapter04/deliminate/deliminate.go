package deliminate

func deliminate(L []string) []string {
	i := 0
	for _, s := range L {
		if i == 0 {
			i++
		} else if s != L[i-1] {
			L[i] = s
			i++
		}
	}
	return L[:i]
}
