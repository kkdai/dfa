package dfa

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
}

func NewDFA() *dfa {
	return &dfa{transition: make(map[transitionInput]int, 0)}
}

func (d *dfa) AddState(state int, isFinal bool) {
	d.totalStates = append(d.totalStates, state)
	if isFinal {
		d.finalStates = append(d.finalStates, state)
	}
}

func (d *dfa) AddTransition(srcState int, input interface{}, dstStateint int) {
	targetTrans := transitionInput{srcState: srcState, input: input}
	d.transition[targetTrans] = dstStateint
}

func (d *dfa) Input(testInput interface{}) bool {
	return false
}
