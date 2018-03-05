package automata

import (
	"errors"
)

//DFA struct
type DFA struct {
	initialState int
	currentState int
	totalStates  []int
	finalStates  []int
	transition   map[transitionData]int
	alphabetMap  map[string]bool
}

//NewDFA creates new dfa
func NewDFA(initialState int, isFinal bool) (*DFA, error) {
	if initialState < 0 {
		return nil, errors.New("Error, estado inicila no puede ser menos que 0")
	}

	newDFA := &DFA{
		initialState: initialState,
		currentState: initialState,
		transition:   make(map[transitionData]int),
		alphabetMap:  make(map[string]bool),
	}

	newDFA.AddState(initialState, isFinal)

	return newDFA, nil
}

//AddState adds new state to dfa
func (dfa *DFA) AddState(state int, isFinal bool) error {
	if state < 0 {
		return errors.New("Error, estado no puede ser menor a 0")
	}

	dfa.totalStates = append(dfa.totalStates, state)

	if isFinal {
		dfa.finalStates = append(dfa.finalStates, state)
	}

	return nil
}

//AddTransition adds transitions from source state to a destination state
func (dfa *DFA) AddTransition(srcState int, alphabet string, dstState int) error {
	srcStateFound := false

	for _, v := range dfa.totalStates {
		if v == srcState {
			srcStateFound = true
		}
	}

	if !srcStateFound {
		return errors.New("Error, no se encontro ningun estado")
	}

	if alphabet == "" {
		return errors.New("Error, estado no puede ser")
	}

	if _, ok := dfa.alphabetMap[alphabet]; !ok {
		dfa.alphabetMap[alphabet] = true
	}

	newTrans := transitionData{
		srcState: srcState,
		alphabet: alphabet,
	}

	dfa.transition[newTrans] = dstState
	return nil
}
