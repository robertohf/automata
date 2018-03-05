package automata

import (
	"errors"
)

//NFA struct implementation
type NFA struct {
	initialState int
	currentState int
	totalStates  []int
	finalStates  []int
	transition   map[transitionData][]int
	alphabetMap  map[string]bool
}

//NewNFA creates new nfa
func NewNFA(initialState int, isFinal bool) *NFA {
	newNFA := &NFA{
		initialState: initialState,
		currentState: initialState,
		transition:   make(map[transitionData][]int),
		alphabetMap:  make(map[string]bool),
	}

	newNFA.AddState(initialState, isFinal)
	return newNFA
}

//AddState adds new state to dfa
func (nfa *NFA) AddState(state int, isFinal bool) error {
	if state < 0 {
		return errors.New("Error, estado no puede ser menor a 0")
	}

	nfa.totalStates = append(nfa.totalStates, state)

	if isFinal {
		nfa.finalStates = append(nfa.finalStates, state)
	}

	return nil
}

//AddTransition adds new transition
func (nfa *NFA) AddTransition(srcState int, alphabet string, dstStates ...int) error {
	srcStateFound := false

	for _, v := range nfa.totalStates {
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

	if _, ok := nfa.alphabetMap[alphabet]; !ok {
		nfa.alphabetMap[alphabet] = true
	}

	var dstArr []int
	for _, destState := range dstStates {
		dstArr = append(dstArr, destState)
	}

	newTrans := transitionData{
		srcState: srcState,
		alphabet: alphabet,
	}

	nfa.transition[newTrans] = dstArr
	return nil
}

//ConvertNFA converts nfa to dfa
func (nfa *NFA) ConvertNFA() (dfa *DFA) {

	dfa = &DFA{
		initialState: nfa.initialState,
		currentState: nfa.currentState,
		transition:   make(map[transitionData]int),
		alphabetMap:  make(map[string]bool),
	}

	var alphabetList []string
	for key := range nfa.alphabetMap {
		alphabetList = append(alphabetList, key)
	}

	for _, state := range nfa.totalStates {
		for _, key := range alphabetList {
			if _, ok := nfa.alphabetMap[key]; ok {
				dfa.alphabetMap[key] = true
			}
			checkTransition := transitionData{srcState: state, alphabet: key}
			if dstState, ok := nfa.transition[checkTransition]; ok {
				if len(dstState) > 1 {
					newState := len(dstState) + 1
					newTrans := transitionData{
						srcState: newState,
						alphabet: key,
					}
					for _, value := range dstState {
						dfa.transition[newTrans] = value
						dfa.transition[checkTransition] = dfa.transition[newTrans]
					}
				} else {
					for _, value := range dstState {
						dfa.transition[checkTransition] = value
					}
				}
			}
		}
	}
	return dfa
}
