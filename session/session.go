package session

import "honnef.co/go/tools/config"

type Session struct {
	//BIG STRUCTURE WITH a LOT of fields
	Config config.Config
}

// If there is no reason to -- Session shouldnt be passed by pointer
// __TODO: Or is it ? I should experiment with it and see what happens
func NewSession() Session {
	return Session{}
}

type InterState struct {
}
type SInterState struct {
	InterState
	ss Session
}

// For middleware purposes
func (s *Session) WithInterState() *SInterState {
	return &SInterState{InterState: InterState{}, ss: *s}
}
