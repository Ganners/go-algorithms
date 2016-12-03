package rebase

import "math"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func posAlphabet(i int) byte {
	return alphabet[i%len(alphabet)]
}

func alphabetPos(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c) - '0'
	}
	if c >= 'A' && c <= 'Z' {
		return (int(c) - 'A') + 10
	}
	if c >= 'a' && c <= 'z' {
		return (int(c) - 'a') + 36
	}
	return 0
}

func Rebase(n int, base int) []byte {
	if base < 2 {
		return nil
	}
	if base > len(alphabet) {
		return nil
	}
	output := make([]byte, 0)
	for n > 0 {
		r := n % base
		n = n / base
		output = append([]byte{posAlphabet(r)}, output...)
	}
	return output
}

func Decimalise(n []byte, base int) int {
	if base > len(alphabet) {
		return 0
	}
	output := 0
	for f, i := len(n)-1, 0; f >= 0; f, i = f-1, i+1 {
		unit := int(math.Pow(float64(base), float64(i)))
		digit := int(alphabetPos(n[f]))
		output += unit * digit
	}
	return output
}
