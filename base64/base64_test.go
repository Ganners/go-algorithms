package base64

import "testing"

func TestBase64Encode(t *testing.T) {
	for _, testCase := range []struct {
		input  string
		output string
	}{
		{
			input:  "Man",
			output: "TWFu",
		},
		{
			input:  "test",
			output: "dGVzdA==",
		},
		{
			input:  "Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure.",
			output: "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=",
		},
	} {
		encoded := Base64Encode(testCase.input)
		if encoded != testCase.output {
			t.Errorf("expected encoded to be %s, got %s", testCase.output, encoded)
		}
	}
}
