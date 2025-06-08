package dataframe

import (
	"path/filepath"

	"github.com/caerbannogwhite/aargh"
)

const (
	NA_TEXT = aargh.NA_TEXT
)

var ctx *aargh.Context
var testDataDir string

func init() {
	ctx = aargh.NewContext()
	testDataDir = filepath.Join("..", "testdata")

	read_G1_1e4_1e2_0_0()
	read_G1_1e5_1e2_0_0()
	read_G1_1e6_1e2_0_0()
	read_G1_1e7_1e2_0_0()
	read_G1_1e4_1e2_10_0()
	read_G1_1e5_1e2_10_0()
	read_G1_1e6_1e2_10_0()
	read_G1_1e7_1e2_10_0()
}
