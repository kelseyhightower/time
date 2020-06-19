package time

import (
	"errors"
	"time"
)

// Juneteenth returns the date Union soldiers arrived in Galveston,
// Texas with news the civil war had ended and the enslaved were now
// free.
//
// Note that this was two and a half years after the Emancipation
// Proclamation in January 1, 1863.
func Juneteenth() time.Time {
	location := time.FixedZone("UTC-6", 0)
	return time.Date(int(1865), time.June, int(19), int(0), int(0), int(0), int(0), location)
}

// CPT returns the time corresponding to being late.
//
// CPT is an American expression referring to a negative racist
// stereotype of African Americans as frequently being late.
//
// CPT always returns an error.
func CPT(u time.Time) (time.Time, error) {
	return time.Time{}, errors.New("time: that's racist")
}
