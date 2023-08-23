package pure

import "testing"

func Test_Prompt_symbol(t *testing.T) {
	want := "❯"
	if got := Prompt_symbol(); got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
