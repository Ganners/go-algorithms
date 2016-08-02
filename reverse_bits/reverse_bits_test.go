package reverse_bits

import (
	"strconv"
	"testing"
)

func TestReverseBits(t *testing.T) {

	for _, test := range []struct {
		input  uint8
		output uint8
	}{
		{
			input:  parseBinary("00000001"),
			output: parseBinary("10000000"),
		},
		{
			input:  parseBinary("00000011"),
			output: parseBinary("11000000"),
		},
		{
			input:  parseBinary("00000111"),
			output: parseBinary("11100000"),
		},
		{
			input:  parseBinary("00001111"),
			output: parseBinary("11110000"),
		},
		{
			input:  parseBinary("11110000"),
			output: parseBinary("00001111"),
		},
		{
			input:  parseBinary("00110000"),
			output: parseBinary("00001100"),
		},
	} {
		if out := ReverseBits(test.input); out != test.output {
			t.Errorf("Output %b does not match expected %b", out, test.output)
		}
		if out := ReverseBitsBitwise(test.input); out != test.output {
			t.Errorf("Output %b does not match expected %b", out, test.output)
		}
		if out := ReverseBitsMagic(test.input); out != test.output {
			t.Errorf("Output %b does not match expected %b", out, test.output)
		}
	}
}

func parseBinary(binary string) uint8 {
	i, err := strconv.ParseInt(binary, 2, 32)
	if err != nil {
		return 0
	}
	return uint8(i)
}

func BenchmarkReverseBits(b *testing.B) {
	input := parseBinary("00110001")
	for i := 0; i < b.N; i++ {
		ReverseBitsMagic(input)
	}
}

func BenchmarkReverseBitsBitwise(b *testing.B) {
	input := parseBinary("00110001")
	for i := 0; i < b.N; i++ {
		ReverseBitsMagic(input)
	}
}

func BenchmarkReverseBitsBitwiseMagic(b *testing.B) {
	input := parseBinary("00110001")
	for i := 0; i < b.N; i++ {
		ReverseBitsMagic(input)
	}
}
