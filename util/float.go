package util

import (
	"fmt"
	"math"
)

func Float64P(i float64) *float64 {
	return &i
}

func FloatToStringWithTruncIfTrailingIsZero(xF float64) string {
	if math.Trunc(xF) == xF {
		return fmt.Sprintf("%.0f", xF)
	} else {
		return fmt.Sprintf("%.2f", xF)
	}
}
