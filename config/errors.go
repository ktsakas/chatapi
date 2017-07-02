package config

import "errors"

var (
	// ErrRecordExists record already exists, happens when we try to insert an entry that already exists
	ErrRecordExists = errors.New("record already exists")
)
