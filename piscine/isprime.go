package piscine

func Isprime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for r := 3; r*r <= nb; r += 2 {
		if nb%r == 0 {
			return false
		}
	}
	return true
}
