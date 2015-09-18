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

	dfa.PrintTransitionTable()
}

func TestAdvanceDFA(t *testing.T) {

	dfa := NewDFA(0, false)
	dfa.AddState(1, true)
	dfa.AddState(2, false)

	dfa.AddTransition(0, "0", 0)
	dfa.AddTransition(0, "1", 1)
	dfa.AddTransition(1, "0", 0)
	dfa.AddTransition(1, "1", 2)
	dfa.AddTransition(2, "0", 2)
	dfa.AddTransition(2, "1", 2)

	dfa.PrintTransitionTable()
	inputs := []string{"0", "0", "1", "0", "1"}

	if !dfa.VerifyInputs(inputs) {
		t.Errorf("Verify inputs is failed")
	}

	//Reset the DFA for another verification
	dfa.Reset()

	//Test go to dead state 2

	inputs2 := []string{"1", "1", "0", "0", "0"}

	if dfa.VerifyInputs(inputs2) {
		t.Errorf("Verify inputs is failed")
	}
}
