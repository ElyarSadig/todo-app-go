package utils

import "testing"

func TestGenerateSecureToken(t *testing.T) {
	testCases := []struct {
		length int
		desc   string
	}{
		{
			length: 10,
			desc:   "20 chars",
		},
		{
			length: 5,
			desc:   "10 chars",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			token, err := GenerateSecureToken(tC.length)
			if err != nil {
				t.Fatal("unexpected error happened:", err)
			}
			if len(token) != tC.length*2 {
				t.Errorf("expected a token with %d chars, but got %d chars", tC.length, len(token))
			}
		})
	}
}
