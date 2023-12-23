package compute

import (
	"fmt"
	"strings"
)

type machineState int

const (
	initialState machineState = iota
	wordState
	spaceState
)

const (
	characterEvent = iota
	spaceEvent
	finiteEvent
)

type transition struct {
	changeState  func()
	appendLetter func(ch rune)
	appendToken  func()
}

type stateMachine struct {
	state       machineState
	transitions map[machineState]map[int]transition
	builder     strings.Builder
	tokens      []string
}

func newStateMachine() *stateMachine {
	sm := &stateMachine{
		state: initialState,
	}

	sm.transitions = map[machineState]map[int]transition{
		initialState: {
			characterEvent: transition{
				appendLetter: sm.appendLetter,
				changeState:  sm.toWordState,
			},
			spaceEvent: transition{
				changeState: sm.toSpaceState,
			},
		},
		wordState: {
			characterEvent: transition{
				appendLetter: sm.appendLetter,
				changeState:  sm.toWordState,
			},
			spaceEvent: transition{
				changeState: sm.toSpaceState,
				appendToken: sm.appendToken,
			},
			finiteEvent: transition{
				appendToken: sm.appendToken,
				changeState: sm.toInitialState,
			},
		},
		spaceState: {
			characterEvent: transition{
				appendLetter: sm.appendLetter,
				changeState:  sm.toWordState,
			},
			spaceEvent: transition{
				changeState: sm.toSpaceState,
			},
			finiteEvent: transition{
				appendToken: sm.appendToken,
				changeState: sm.toInitialState,
			},
		},
	}

	return sm
}

func (sm *stateMachine) parse(query string) ([]string, error) {
	for _, ch := range query {
		if isLetter(ch) {
			sm.tick(characterEvent, ch)
		} else if isWhiteSpace(ch) {
			sm.tick(spaceEvent, ch)
		} else {
			return nil, fmt.Errorf("invalid character %c", ch)
		}
	}

	sm.tick(finiteEvent, ' ')

	return sm.tokens, nil
}

func (sm *stateMachine) tick(event int, ch rune) {
	transition := sm.transitions[sm.state][event]

	if transition.changeState != nil {
		transition.changeState()
	}

	if transition.appendLetter != nil {
		transition.appendLetter(ch)
	}

	if transition.appendToken != nil {
		transition.appendToken()
	}
}

func (sm *stateMachine) appendLetter(ch rune) {
	sm.builder.WriteRune(ch)
}

func (sm *stateMachine) toWordState() {
	sm.state = wordState
}

func (sm *stateMachine) toSpaceState() {
	sm.state = spaceState
}

func (sm *stateMachine) toInitialState() {
	sm.state = initialState
}

func (sm *stateMachine) appendToken() {
	sm.tokens = append(sm.tokens, sm.builder.String())
	sm.builder.Reset()
}

func isWhiteSpace(ch rune) bool {
	return ch == '\t' || ch == '\n' || ch == ' '
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '*') ||
		(ch == '/') ||
		(ch == '_')
}
