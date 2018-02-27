package automata

import (
  "errors"
)

type DFA struct {
	initialState	int
	currentState	int
	totalStates		[]int
	finalStates		[]int
	transition		map[transitionData]int
}

type transitionData struct {
  srcState      int
  alphabet      string  
} 

func NewDFA(initialState int, isFinal bool) *DFA {
	newDFA := &DFA{
		initialState: initialState,
		currentState: initialState,
    transition: make(map[transitionData]int),
  }
  
  newDFA.AddState(initialState, isFinal)

  return newDFA
}

func (dfa *DFA) AddState(state int, isFinal bool) error {
  if state < 0 {
    return errors.New("Error, estado no puede ser menor a 0.")
  }

  dfa.totalStates = append(dfa.totalStates, state)

  if isFinal {
    dfa.finalStates = append(dfa.finalStates, state)
  }

  return nil
}

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
  
  newTrans := transitionData {
    srcState: srcState,
    alphabet: alphabet,
  }

  dfa.transition[newTrans] = dstState

  return nil
}

