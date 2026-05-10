package Delegates

import (
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

// user implementations
type ConcernsImplementation struct {
	Concern
}

func (s *ConcernsImplementation) Scope() { s.Where("id", 2) }
