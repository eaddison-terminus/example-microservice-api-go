package repo

import "errors"

// ErrNot Found is required when no record is found.

var ErrNotFound = errors.New("not found")
