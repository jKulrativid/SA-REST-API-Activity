package main

import "errors"

var ErrEntityNotFound = errors.New("entity not found")
var ErrConflict = errors.New("conflict")
