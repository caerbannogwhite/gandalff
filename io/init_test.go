package io

import (
	"github.com/caerbannogwhite/gandalff"
)

const (
	NA_TEXT = gandalff.NA_TEXT
)

var ctx *gandalff.Context

func init() {
	ctx = gandalff.NewContext()
}
