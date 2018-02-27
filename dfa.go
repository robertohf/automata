package automata

import (
  "errors"
)

type DFA struct {
	initialState	int
	currentState	int
	totalStates		[]int
	finalStates		[]int
  transition		map[dfaTransitionData]int
  symbolMap     map[string]bool
}

type dfaTransitionData struct {
  srcState      int
  symbol        string  
} 

func NewDFA(initialState int, isFinal bool) *DFA {
	newDFA := &DFA{
		initialState: initialState,
		currentState: initialState,
    transition: make(map[dfaTransitionData]int),
    symbolMap:  make(map[string]bool),
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

func (dfa *DFA) AddTransition(srcState int, symbol string, dstState int) error {
  srcStateFound := false

  for _, v := range dfa.totalStates {
    if v == srcState {
      srcStateFound = true
    }
  }

  if !srcStateFound {
    return errors.New("Error, no se encontro ningun estado")
  }
  
  if _, ok := dfa.symbolMap[symbol]; !ok {
    dfa.symbolMap[symbol] = true
  }

  newTrans := dfaTransitionData {
    srcState: srcState,
    symbol: symbol,
  }

  dfa.transition[newTrans] = dstState

  return nil
}

func (dfa *DFA) ValidateFinalState() bool {
  for _, val := dfa.finalState {
    if v == dfa.currentState {
      return true
    }
  }
  return false
}

func (dfa *DFA) ValidateSymbol(symbols []string) bool {
  for _, val := range symbols {
    
  }
  return dfa.ValidateFinalState()
}