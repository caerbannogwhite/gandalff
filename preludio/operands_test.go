package preludio

import (
	"testing"
)

var be *ByteEater

func init() {
	be = new(ByteEater).InitVM()
}

func TestMul(t *testing.T) {

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL
	{
		// BOOL * BOOL
		b1.appendOperand(OP_BINARY_MUL, b2)
		b1.solve(be)

		if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
				t.Error("Expected [0, 0, 1, 0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * INTEGER
		b1.appendOperand(OP_BINARY_MUL, in)
		b1.solve(be)

		if !b1.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := b1.getIntegerVector()
			if v[0] != 1 || v[1] != 0 || v[2] != 3 || v[3] != 0 {
				t.Error("Expected [1, 0, 3, 0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * FLOAT
		b1.appendOperand(OP_BINARY_MUL, fl)
		b1.solve(be)

		if !b1.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := b1.getFloatVector()
			if v[0] != 5.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 0.0 {
				t.Error("Expected [5.0, 0.0, 7.0, 0.0]")
			}
		}

		// reset b1
		b1 = newPInternTerm([]bool{true, false, true, false})

		// BOOL * STRING
		b1.appendOperand(OP_BINARY_MUL, st)
		b1.solve(be)

		if !b1.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := b1.getStringVector()
			if v[0] != "a" || v[1] != "" || v[2] != "c" || v[3] != "" {
				t.Error("Expected [\"a\", \"\", \"c\", \"\"]")
			}
		}
	}

	// INTEGER
	{
		// INTEGER * BOOL
		in.appendOperand(OP_BINARY_MUL, b2)
		in.solve(be)

		if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 0 || v[1] != 0 || v[2] != 3 || v[3] != 4 {
				t.Error("Expected [0, 0, 3, 4]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * INTEGER
		in.appendOperand(OP_BINARY_MUL, in)
		in.solve(be)

		if !in.isIntegerVector() {
			t.Error("Expected integer vector type")
		} else {
			v, _ := in.getIntegerVector()
			if v[0] != 1 || v[1] != 4 || v[2] != 9 || v[3] != 16 {
				t.Error("Expected [1, 4, 9, 16]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * FLOAT
		in.appendOperand(OP_BINARY_MUL, fl)
		in.solve(be)

		if !in.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := in.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [5.0, 12.0, 21.0, 32.0]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})

		// INTEGER * STRING
		in.appendOperand(OP_BINARY_MUL, st)
		in.solve(be)

		if !in.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := in.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset in
		in = newPInternTerm([]int{1, 2, 3, 4})
	}

	// FLOAT
	{
		// FLOAT * BOOL
		fl.appendOperand(OP_BINARY_MUL, b2)
		fl.solve(be)

		if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 0.0 || v[1] != 0.0 || v[2] != 7.0 || v[3] != 8.0 {
				t.Error("Expected [0.0, 0.0, 7.0, 8.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * INTEGER
		fl.appendOperand(OP_BINARY_MUL, in)
		fl.solve(be)

		if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 5.0 || v[1] != 12.0 || v[2] != 21.0 || v[3] != 32.0 {
				t.Error("Expected [0.0, 0.0, 21.0, 0.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * FLOAT
		fl.appendOperand(OP_BINARY_MUL, fl)
		fl.solve(be)

		if !fl.isFloatVector() {
			t.Error("Expected float vector type")
		} else {
			v, _ := fl.getFloatVector()
			if v[0] != 25.0 || v[1] != 36.0 || v[2] != 49.0 || v[3] != 64.0 {
				t.Error("Expected [25.0, 36.0, 49.0, 64.0]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})

		// FLOAT * STRING
		fl.appendOperand(OP_BINARY_MUL, st)
		fl.solve(be)

		if !fl.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := fl.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset fl
		fl = newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	}

	// STRING
	{
		// STRING * BOOL
		st.appendOperand(OP_BINARY_MUL, b2)
		st.solve(be)

		if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "" || v[1] != "" || v[2] != "ccc" || v[3] != "" {
				t.Error("Expected [\"\", \"\", \"ccc\", \"\"]")
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * INTEGER
		st.appendOperand(OP_BINARY_MUL, in)
		st.solve(be)

		if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * FLOAT
		st.appendOperand(OP_BINARY_MUL, fl)
		st.solve(be)

		if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "ccc" || v[3] != "dddd" {
				t.Error("Expected [\"a\", \"bb\", \"ccc\", \"dddd\"]")
			}
		}

		// reset st
		st = newPInternTerm([]string{"a", "b", "c", "d"})

		// STRING * STRING
		st.appendOperand(OP_BINARY_MUL, st)
		st.solve(be)

		if !st.isStringVector() {
			t.Error("Expected string vector type")
		} else {
			v, _ := st.getStringVector()
			if v[0] != "a" || v[1] != "bb" || v[2] != "cccc" || v[3] != "dddddd" {
				t.Error("Expected [\"a\", \"bb\", \"cccc\", \"dddddd\"]")
			}
		}
	}
}

func TestDiv(t *testing.T) {

	b1 := newPInternTerm([]bool{true, false, true, false})
	b2 := newPInternTerm([]bool{false, false, true, true})
	in := newPInternTerm([]int{1, 2, 3, 4})
	fl := newPInternTerm([]float64{5.0, 6.0, 7.0, 8.0})
	st := newPInternTerm([]string{"a", "b", "c", "d"})

	// BOOL / BOOL
	b1.appendOperand(OP_BINARY_DIV, b2)
	b1.solve(be)

	if !b1.isIntegerVector() {
		t.Error("Expected integer vector type")
	} else {
		v, _ := b1.getIntegerVector()
		if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
			t.Error("Expected [0, 0, 1, 0]")
		}
	}

	// reset b1
	b1 = newPInternTerm([]bool{true, false, true, false})

	// BOOL / INTEGER
	b1.appendOperand(OP_BINARY_DIV, in)
	b1.solve(be)

	if !b1.isIntegerVector() {
		t.Error("Expected integer vector type")
	} else {
		v, _ := b1.getIntegerVector()
		if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
			t.Error("Expected [0, 0, 1, 0]")
		}
	}

	// reset b1
	b1 = newPInternTerm([]bool{true, false, true, false})

	// BOOL / FLOAT
	b1.appendOperand(OP_BINARY_DIV, fl)
	b1.solve(be)

	if !b1.isIntegerVector() {
		t.Error("Expected integer vector type")
	} else {
		v, _ := b1.getIntegerVector()
		if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
			t.Error("Expected [0, 0, 1, 0]")
		}
	}

	// reset b1
	b1 = newPInternTerm([]bool{true, false, true, false})

	// BOOL / STRING
	b1.appendOperand(OP_BINARY_DIV, st)
	b1.solve(be)

	if !b1.isIntegerVector() {
		t.Error("Expected integer vector type")
	} else {
		v, _ := b1.getIntegerVector()
		if v[0] != 0 || v[1] != 0 || v[2] != 1 || v[3] != 0 {
			t.Error("Expected [0, 0, 1, 0]")
		}
	}

	// reset b1
	b1 = newPInternTerm([]bool{true, false, true, false})

	// INTEGER / BOOL
	in.appendOperand(OP_BINARY_DIV, b1)
	in.solve(be)

	if !in.isIntegerVector() {
		t.Error("Expected integer vector type")
	} else {
		v, _ := in.getIntegerVector()
		if v[0] != 0 || v[1] != 0 || v[2] != 3 || v[3] != 0 {
			t.Error("Expected [0, 0, 3, 0]")
		}
	}
}
