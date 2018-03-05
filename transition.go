package automata

type transitionData struct {
	srcState int
	alphabet string
}

//DFAValidateFinalState compares current state to final state
func (dfa *DFA) DFAValidateFinalState() bool {
	for _, val := range dfa.finalStates {
		if val == dfa.currentState {
			return true
		}
	}
	return false
}

//DFAValidateAlphabet DFA using input symbols
func (dfa *DFA) DFAValidateAlphabet(alphabet []string) bool {
	for _, value := range alphabet {
		alphabetTransition := transitionData{
			srcState: dfa.currentState,
			alphabet: value,
		}
		if value, ok := dfa.transition[alphabetTransition]; ok {
			dfa.currentState = value
		}
	}
	return dfa.DFAValidateFinalState()
}

//DFAReset resets current staet to init
func (dfa *DFA) DFAReset() {
	dfa.currentState = dfa.initialState
}

//NFAValidateFinalState compares current state to final state
func (nfa *NFA) NFAValidateFinalState() bool {
	for _, value := range nfa.finalStates {
		if value == nfa.currentState {
			return true
		}
	}
	return false
}

//NFAValidateAlphabet DFA using input symbols
func (nfa *NFA) NFAValidateAlphabet(alphabet []string) bool {

	for _, value := range alphabet {
		var updateCurrentState int

		alphabetTransition := transitionData{
			srcState: nfa.currentState,
			alphabet: value,
		}
		if val, ok := nfa.transition[alphabetTransition]; ok {
			for dstState := range val {
				updateCurrentState = dstState
			}
		}
		nfa.currentState = updateCurrentState
	}

	return nfa.NFAValidateFinalState()
}

//NFAReset resets current staet to init
func (nfa *NFA) NFAReset() {
	nfa.currentState = nfa.initialState
}
