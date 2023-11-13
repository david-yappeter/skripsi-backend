package number_speller

import (
	"fmt"
	"strings"
)

var (
	idNumberSpeller Interface = newNumberSpeller("id", integerToIDID)
)

func integerToIDID(input int) string {
	var indonesianMegas = []string{"", "ribu", "juta", "milyar", "triliun", "kuadriliun", "kuintiliun", "sekstiliun", "septiliun", "oktiliun", "noniliun", "desiliun", "undesiliun", "duodesiliun", "tredesiliun", "kuatuordesiliun"}
	var indonesianUnits = []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	var indonesianTens = []string{"", "sepuluh", "dua puluh", "tiga puluh", "empat puluh", "lima puluh", "enam puluh", "tujuh puluh", "delapan puluh", "sembilan puluh"}
	var indonesianTeens = []string{"sepuluh", "sebelas", "dua belas", "tiga belas", "empat belas", "lima belas", "enam belas", "tujuh belas", "delapan belas", "sembilan belas"}

	words := []string{}

	if input < 0 {
		words = append(words, "minus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)

	// zero is a special case
	if len(triplets) == 0 {
		return "nol"
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10

		if hundreds == 1 {
			words = append(words, "seratus")
		} else if hundreds > 0 {
			words = append(words, indonesianUnits[hundreds], "ratus")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, indonesianUnits[units])
		case 1:
			words = append(words, indonesianTeens[units])
		default:
			if units > 0 {
				word := fmt.Sprintf("%s %s", indonesianTens[tens], indonesianUnits[units])
				words = append(words, word)
			} else {
				words = append(words, indonesianTens[tens])
			}
		}

	tripletEnd:
		// mega
		if mega := indonesianMegas[idx]; mega != "" {
			// exception for 1000
			if idx == 1 && triplet == 1 {
				words = append(words[0:len(words)-1], "seribu")
			} else {
				words = append(words, mega)
			}
		}
	}

	return strings.Join(words, " ")
}
