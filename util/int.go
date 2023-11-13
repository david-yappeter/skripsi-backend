package util

import (
	"math"
	"strconv"
)

var numByRoman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var romanByNum = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var romanByNumMaxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func IntP(i int) *int {
	return &i
}

func MustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}

func RomanToInt(n string) int {
	out := 0

	ln := len(n)
	for i := 0; i < ln; i++ {
		c := string(n[i])

		vc := numByRoman[c]
		if i < ln-1 {
			cnext := string(n[i+1])
			vcnext := numByRoman[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}

	return out
}

func IntToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += romanByNum[v]
		n -= v
	}

	return out
}

func highestDecimal(n int) int {
	for _, v := range romanByNumMaxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func ArrayUint8ToArrayInt16(arr []uint8) []int16 {
	numbers := []int16{}

	for _, v := range arr {
		numbers = append(numbers, int16(v))
	}

	return numbers
}
