package io

import (
	"path/filepath"

	"github.com/caerbannogwhite/gandalff"
)

const (
	NA_TEXT = gandalff.NA_TEXT
)

var ctx *gandalff.Context
var testDataFolder = filepath.Join("..", "testdata")

func init() {
	ctx = gandalff.NewContext()
}
