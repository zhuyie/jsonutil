package jsonutil

import "testing"

func TestBasic(t *testing.T) {
	var parsed interface{}
	var err error

	err = Unmarshal([]byte("{}"), &parsed)
	if err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	err = Unmarshal([]byte("{\"abd\":\n\"xyz\"}"), &parsed)
	if err != nil {
		t.Errorf("Unmarshal error: %v", err)
	}

	err = Unmarshal([]byte("{\"abd\":\n\"xyz\",}"), &parsed)
	if err != nil {
		syntax, ok := err.(*SyntaxError)
		if !ok {
			t.Errorf("Expect SyntaxError, got %v", err)
		}
		if syntax.Line != 2 {
			t.Errorf("Line should be 2, got %v", syntax.Line)
		}
		if syntax.Pos != 8 {
			t.Errorf("Pos should be 8, got %v", syntax.Pos)
		}
	} else {
		t.Error("Expect SyntaxError, got nil")
	}
}
