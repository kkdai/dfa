package dfa

import "testing"

func TestBasic(t *testing.T) {
	dfa := NewDFA(0, false)
	dfa.AddState(1, false)
	dfa.AddState(2, true)

	dfa.AddTransition(0, "a", 1)
	dfa.AddTransition(1, "b", 2)

	if ret := dfa.Input("a"); ret != 1 {
		t.Errorf("Expect 1, but get %d\n", ret)
	}

	if ret := dfa.Input("b"); ret != 2 {
		t.Errorf("Expect 2, but get %d\n", ret)
	}

	if !dfa.Verify() {
		t.Errorf("Verify is failed")
	}
}

func TestVerifyInputs(t *testing.T) {
	dfa := NewDFA(0, false)
	dfa.AddState(1, false)
	dfa.AddState(2, true)

	dfa.AddTransition(0, "a", 1)
	dfa.AddTransition(1, "b", 2)

	var inputs []string
	inputs = append(inputs, "a")
	inputs = append(inputs, "b")
	if !dfa.VerifyInputs(inputs) {
		t.Errorf("Verify Inputs is failed")
	}
}
