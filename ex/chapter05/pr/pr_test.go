package pr

import "testing"

func TestPr(t *testing.T) {

	if pr() != 0xdeadbeef {
		t.Error("expected so see dead beef")
	}

}
