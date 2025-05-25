package series

import (
	"testing"

	"github.com/caerbannogwhite/gandalff"
)

func Test_Series(t *testing.T) {

	s := NewSeries([]bool{true, false, true, false, true, false, true, false, true, false}, nil, true, false, ctx)

	r := s.Append(true).
		Append([]gandalff.NullableBool{{Valid: true, Value: true}, {Valid: true, Value: false}}).
		Filter([]bool{true, false, true, false, true, false, true, false, true, false, true, true, false})

	if e, ok := r.(Errors); ok {
		t.Errorf("Expected a series, got an error: %s", e.GetError())
	}

	if r.Len() != 7 {
		t.Errorf("Expected length of 7, got %d", r.Len())
	}
}
