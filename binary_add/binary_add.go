package binary_add

// AddBinary adds two slices of binary values (int slices, meant to be supplied
// with 0s and 1s)
func AddBinary(a []int, b []int) []int {

	c := make([]int, len(a)+1)

	carry := 0
	i := 0

	for ; i < len(a); i++ {
		c[i] = (a[i] + b[i] + carry) % 2
		carry = (a[i] + b[i] + carry) / 2
	}
	c[i] = carry

	return c
}
