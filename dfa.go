package dfa

type dfa struct {
	initState    int
	currentState int
	totalStates  []int
	finalStates  []int
	transition   map[int]int
}

func NewDFA() *dfa {
	return &dfa{}
}
