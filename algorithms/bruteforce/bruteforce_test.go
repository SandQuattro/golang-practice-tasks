package bruteforce

import "testing"

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"0",
		"12",
		"333",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}
