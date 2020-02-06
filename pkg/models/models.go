package models

import (
	"errors"
	"time"
)

// ErrNoRecord : error for no record match
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet : struct defining the snippet table model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
