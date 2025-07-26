package state

import (
	"fmt"
	"slices"

	. "github.com/binje/Fields/actions"
)

// For benchmarking: nextState is now a slice, indexed by Action (assumes Action is a dense int)
type State struct {
	finished  bool
	nextState []*State
	prev      *State
	vp        int
}

// NewState creates a new state node, optionally linking to a previous state.
func newState(prev *State) *State {
	return &State{
		finished:  false,
		nextState: nil,
		prev:      prev,
		vp:        -10000,
	}
}

func Root() *State {
	return newState(nil)
}

func (s *State) LoadActions(actions []Action) {
	if len(s.nextState) != 0 {
		return
	}
	
	maxAction := slices.MaxFunc(actions, func(x, y Action) int {
		return int(x) - int(y)
	})
	// Allocate nextState with enough space for all possible actions
	s.nextState = make([]*State, int(maxAction)+1)
	for _, action := range actions {
		s.nextState[int(action)] = newState(s)
	}
}

func (s *State) TakeAction(a Action) *State {
	if int(a) >= len(s.nextState) || s.nextState[int(a)] == nil {
		panic(a)
	}
	return s.nextState[int(a)]
}

func (s *State) IsFinished(a Action) bool {
	if int(a) >= len(s.nextState) || s.nextState[int(a)] == nil {
		return false
	}
	return s.nextState[int(a)].finished
}

func (s *State) RootFinished() bool {
	return s.finished
}

func (s *State) MarkFinished(vp int) {
	s.finished = true
	s.vp = vp
	s.walkBack()
}

func (s *State) walkBack() {
	s = s.prev
	if s == nil {
		return
	}
	for _, ns := range s.nextState {
		if ns != nil && !ns.finished {
			return
		}
	}
	s.finished = true
	for _, ns := range s.nextState {
		if ns != nil && ns.vp > s.vp {
			s.vp = ns.vp
		}
	}
	s.walkBack()
}

func (s *State) PrintCurrNext() {
	fmt.Println()
	fmt.Println("PRINTING STATE")
	fmt.Println(s)
	fmt.Println()

	for k, v := range s.nextState {
		if v != nil {
			fmt.Println(k, v)
			fmt.Println()
		}
	}
}

func (s *State) GetVp() int {
	return s.vp
}
