package state

import (
	"fmt"

	. "github.com/binje/Fields/actions"
)

type State struct {
	finished  bool
	nextState map[Action]*State
	prev      *State
}

func Root() *State {
	return newState(nil)
}

func (s *State) LoadActions(a []Action) {
	if len(s.nextState) != 0 {
		return
	}
	for _, act := range a {
		s.nextState[act] = newState(s)
	}
}

func (s *State) TakeAction(a Action) *State {
	ns, ok := s.nextState[a]
	if !ok {
		panic(a)
	}
	return ns
}

func newState(s *State) *State {
	return &State{
		false,
		make(map[Action]*State),
		s,
	}
}

func (s *State) IsFinished(a Action) bool {
	fmt.Println("looking for: ", a)
	for k, v := range s.nextState {
		fmt.Println(k, v)
	}
	return s.nextState[a].finished
}

func (s *State) RootFinished() bool {
	return s.finished
}

func (s *State) MarkFinished() {
	s.finished = true
	s.walkBack()
}

func (s *State) walkBack() {
	s = s.prev
	if s == nil {
		return
	}
	for _, ns := range s.nextState {
		// there is an unfinished option
		if !ns.finished {
			return
		}
	}
	// all next states are finished
	s.finished = true
	s.walkBack()
}
