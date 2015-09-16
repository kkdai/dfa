package dfa

import "fmt"

type transitionInput struct {
	srcState int
	input    interface{}
}

type dfa struct {
	initState    int
	currentState int
	totalStates  []int
	finalStates  []int
	transition   map[transitionInput]int
	inputMap     map[interface{}]bool
}

//New a new DFA
func NewDFA(initState int, isFinal bool) *dfa {
	retDFA := &dfa{
		transition:   make(map[transitionInput]int),
		inputMap:     make(map[interface{}]bool),
		initState:    initState,
		currentState: initState}

	retDFA.AddState(initState, isFinal)
	return retDFA
}

//Add new state in this DFA
func (d *dfa) AddState(state int, isFinal bool) {
	if state == -1 {
		fmt.Println("Cannot add state as -1, it is dead state")
		return
	}

	d.totalStates = append(d.totalStates, state)
	if isFinal {
		d.finalStates = append(d.finalStates, state)
	}
}

//Add new transition function into DFA
func (d *dfa) AddTransition(srcState int, input interface{}, dstStateint int) {
	find := false

	for _, v := range d.totalStates {
		if v == srcState {
			find = true
		}
	}

	if !find {
		fmt.Println("No such state:", srcState, " in current DFA")
		return
	}

	//find input if exist in DFA input List
	if _, ok := d.inputMap[input]; !ok {
		//not exist, new input in this DFA
		d.inputMap[input] = true
	}

	targetTrans := transitionInput{srcState: srcState, input: input}
	d.transition[targetTrans] = dstStateint
}

func (d *dfa) Input(testInput interface{}) int {
	intputTrans := transitionInput{srcState: d.currentState, input: testInput}
	if val, ok := d.transition[intputTrans]; ok {
		d.currentState = val
		return val
	} else {
		return -1 //dead state
	}
}

func (d *dfa) verify() bool {
	for _, val := range d.finalStates {
		if val == d.currentState {
			return true
		}
	}
	return false
}

//Reset DFA state to initilize state, but all state and transition function will remain
func (d *dfa) Reset() {
	d.currentState = d.initState
}

//Verify if list of input could be accept by DFA or not
func (d *dfa) VerifyInputs(inputs []interface{}) bool {
	for _, v := range inputs {
		d.Input(v)
	}
	return d.verify()
}

//To print detail transition table contain of current DFA
func (d *dfa) PrintTransitionTable() {
	fmt.Println("===================================================")
	for key, val := range d.transition {

		fmt.Printf("%d |\t %d \n", key.srcState, val)
	}
}
