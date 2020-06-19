package time

import (
	"time"
	"testing"

)

func TestCPT(t *testing.T) {
	_, err := CPT(time.Now())
	if err == nil {
		t.Fatalf("CPT should always return an error")
	}
}

func TestJuneteenth(t *testing.T) {
	date := Juneteenth()
	if date.Month() != time.June {
		t.Fatalf("want %v, got %v", time.June, date.Month())
	}
}
