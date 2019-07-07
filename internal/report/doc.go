// Code reused from Golang present tool:
// Copyright 2011 The Go Authors. All rights reserved.

package report

import (
	"time"
)

// Doc represents an entire document.
type Doc struct {
	Title      string
	Subtitle   string
	Time       time.Time
	Authors    []Author
	TitleNotes []string
	Sections   []Section
	Tags       []string
}

// Author represents the person who wrote and/or is presenting the document.
type Author struct {
	Elem []Elem
}

// Section represents a section of a document (such as a presentation slide)
// comprising a title and a list of elements.
type Section struct {
	Number  []int
	Title   string
	Elem    []Elem
	Notes   []string
	Classes []string
	Styles  []string
}

// Elem defines the interface for a present element. That is, something that
// can provide the name of the template used to render the element.
type Elem interface {
	TemplateName() string
}

// Text represents an optionally preformatted paragraph.
type Text struct {
	Lines []string
	Pre   bool
}

// List represents a bulleted list.
type List struct {
	Bullet []string
}

//nolint
// Lines is a helper for parsing line-based input.
type Lines struct {
	line int // 0 indexed, so has 1-indexed number of last line returned
	text []string
}
