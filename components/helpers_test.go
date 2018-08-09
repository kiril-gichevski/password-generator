package components

import (
	"testing"
)

func TestConvertToInt(t *testing.T) {
	total := ConvertToInt("25")
	if total != 25 {
		t.Errorf("Converting to int failed")
	}
}
