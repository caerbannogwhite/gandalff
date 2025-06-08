package io

import (
	"path/filepath"

	"github.com/caerbannogwhite/aargh"
)

const (
	NA_TEXT = aargh.NA_TEXT
)

var ctx *aargh.Context
var testDataFolder = filepath.Join("..", "testdata")

func init() {
	ctx = aargh.NewContext()
}
