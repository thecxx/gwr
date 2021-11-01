package gwr

import "testing"

func TestScan(t *testing.T) {
	lex, err := Scan("{header.from_uri: /^[0-9]+$/.match($0), @default: \"https://www.baidu.com\", @error: \"invalid parameter\"}")
	if err != nil {
		t.Errorf("error: %s\n", err)
	} else {
		for key, value := range lex.tokens {
			t.Logf("<%d, %d, %s>\n", key, value.Type, value.Value)
		}
	}
}
