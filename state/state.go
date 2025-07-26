package state

import (
	"fmt"

	. "github.com/binje/Fields/actions"
)

// For benchmarking: nextState is now a slice, indexed by Action (assumes Action is a dense int)
type State struct {
	finished  bool
	nextState []*State
	prev      *State
	vp        int
}

func Root() *State {
	return newState(nil)
}

func (s *State) LoadActions(a []Action) {
	if len(s.nextState) != 0 {
		return
	}
	maxAction := 0
	for _, act := range a {
		if int(act) > maxAction {
			maxAction = int(act)
		}
	}
	s.nextState = make([]*State, maxAction+1)
	for _, act := range a {
		s.nextState[int(act)] = newState(s)
	}
}

func (s *State) TakeAction(a Action) *State {
	if int(a) >= len(s.nextState) || s.nextState[int(a)] == nil {
		panic(a)
	}
	return s.nextState[int(a)]
}

func newState(s *State) *State {
	return &State{
		false,
		nil,
		s,
		-10000,
	}
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
