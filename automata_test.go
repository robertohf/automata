package automata

import (
	"fmt"
	"testing"
)

//TestAutomata testing our dfa
func TestAutomata(t *testing.T) {
	dfa, err := NewDFA(0, false)
	if err != nil {
		t.Fatal(err)
	}
	dfa.AddState(1, false)
	dfa.AddState(2, true)
	dfa.AddTransition(0, "a", 1)
	dfa.AddTransition(1, "a", 1)
	dfa.AddTransition(1, "b", 2)
	dfa.AddTransition(0, "b", 2)
	dfa.AddTransition(2, "a", 2)

	var symbolList []string
	for key := range dfa.alphabetMap {
		fmt.Printf("\t %s|", key)
		symbolList = append(symbolList, key)
	}

	fmt.Printf("\n\n")

	for _, state := range dfa.totalStates {
		fmt.Printf("%d |", state)
		for _, key := range symbolList {
			checkInput := transitionData{srcState: state, alphabet: key}
			if dstState, ok := dfa.transition[checkInput]; ok {
				fmt.Printf("\t %d|", dstState)
			} else {
				fmt.Printf("\t θ|")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	nfa := NewNFA(0, true)
	nfa.AddState(1, true)
	nfa.AddState(2, false)
	nfa.AddState(3, true)

	nfa.AddTransition(0, "a", 1, 2)
	nfa.AddTransition(2, "b", 3)
	nfa.AddTransition(3, "a", 2)
	fmt.Println(nfa)

	var alphabetList []string
	for key := range dfa.alphabetMap {
		fmt.Printf("\t %s|", key)
		alphabetList = append(alphabetList, key)
	}

	fmt.Printf("\n\n")

	for _, state := range nfa.totalStates {
		fmt.Printf("%d |", state)
		for _, key := range alphabetList {
			checkInput := transitionData{srcState: state, alphabet: key}
			if dstState, ok := nfa.transition[checkInput]; ok {
				fmt.Printf("\t %d|", dstState)
			} else {
				fmt.Printf("\t θ|")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	fmt.Println(nfa.ConvertNFA())

	fmt.Printf("\n")
}

