package solutions

// Gcd find the greatest common denominator
func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
}
