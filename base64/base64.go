package base64

var (
	table = []byte{}
)

func init() {
	for b := 'A'; b <= 'Z'; b++ {
		table = append(table, byte(b))
	}
	for b := 'a'; b <= 'z'; b++ {
		table = append(table, byte(b))
	}
	for b := '0'; b <= '9'; b++ {
		table = append(table, byte(b))
	}
	table = append(table, '+', '/')
}

// Base64Encode will encode a string into a base64 encoded string
func Base64Encode(in string) string {
	encoded := ""
	word := uint(0)
	for p, c := range in {
		word |= uint(c) << uint(8*(2-p%3))
		blockEnd := p%3 == 2
		bytesEnd := p >= len(in)-1
		if blockEnd || bytesEnd {
			for i := 18; i >= 0; i -= 6 {
				b := table[((63<<uint(i))&word)>>uint(i)]
				if b == 'A' && bytesEnd && i <= 6 {
					b = '='
				}
				encoded += string(b)
			}
			word = 0
		}
	}
	return encoded
}
