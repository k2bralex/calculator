package pkg

func Contains(r rune, sl []rune) bool {
	for _, v := range sl {
		if v == r {
			return true
		}
	}
	return false
}
