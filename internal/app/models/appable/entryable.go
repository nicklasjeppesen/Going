package Delegates

import (
	. "github.com/nicklasjeppesen/going_internal/super/db/types"
)

// The interface to be implemented by models that can be entryable
type IEntry interface {
	IRepository

	//
	// pre implemented function
	PrintHi() string

	//
	//Abstract function
	//needed to implemnt delegatable models
	//
	Title() string

	// interface method
	CustomID() string
}

// The struct that can be embedded in models to make them delegatable entries
type EntryAble struct {
}

func (e *EntryAble) PrintHi() string {
	return "Hi from EntryAble"
}
