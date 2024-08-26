package tst

import "testing"

func TestHello(t *testing.T) {
	for _, want := range []string{"Hello, world", "Hello, World!", "Hello, Go!"} {
		t.Run(want, func(t *testing.T) {
			got := Hello()
			if got != want {
				t.Errorf("got %q want %q", got, want)
				t.Fail()
			}
		})
	}

}
