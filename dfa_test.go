package dfa

import "testing"

func TestBasic(t *testing.T) {
	dfa := NewDFA(0, false)
	dfa.AddState(1, false)
	dfa.AddState(2, true)

}
