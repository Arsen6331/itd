package translit

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// https://en.wikipedia.org/wiki/Hangul_Jamo_%28Unicode_block%29
var jamoBlock = &unicode.RangeTable{
	R16: []unicode.Range16{{
		Lo:     0x1100,
		Hi:     0x11FF,
		Stride: 1,
	}},
}

// https://en.wikipedia.org/wiki/Hangul_Syllables
var syllablesBlock = &unicode.RangeTable{
	R16: []unicode.Range16{{
		Lo:     0xAC00,
		Hi:     0xD7A3,
		Stride: 1,
	}},
}

// https://en.wikipedia.org/wiki/Hangul_Compatibility_Jamo
var compatJamoBlock = &unicode.RangeTable{
	R16: []unicode.Range16{{
		Lo:     0x3131,
		Hi:     0x318E,
		Stride: 1,
	}},
}

// KoreanTranslit implements transliteration for Korean.
//
// This was translated to Go from the code in https://codeberg.org/Freeyourgadget/Gadgetbridge
type KoreanTranslit struct{}

func (KoreanTranslit) Init() {}

// User input consisting of isolated jamo is usually mapped to the KS X 1001 compatibility
// block, but jamo resulting from decomposed syllables are mapped to the modern one. This
// function maps compat jamo to modern ones where possible and returns all other characters
// unmodified.
//
// https://en.wikipedia.org/wiki/Hangul_Compatibility_Jamo
// https://en.wikipedia.org/wiki/Hangul_Jamo_%28Unicode_block%29
func decompatJamo(jamo rune) rune {
	// KS X 1001 Hangul filler, not used in modern Unicode. A useful landmark in the
	// compatibility jamo block.
	// https://en.wikipedia.org/wiki/KS_X_1001#Hangul_Filler
	var hangulFiller rune = 0x3164

	// Ignore characters outside compatibility jamo block
	if !unicode.In(jamo, compatJamoBlock) {
		return jamo
	}

	// Vowels are contiguous, in the same order, and unambiguous so it's a simple offset.
	if jamo >= 0x314F && jamo < hangulFiller {
		return jamo - 0x1FEE
	}

	// Consonants are organized differently. No clean way to do this.
	// The compatibility jamo block doesn't distinguish between Choseong (leading) and Jongseong
	// (final) positions, but the modern block does. We map to Choseong here.
	switch jamo {
	case 0x3131:
		return 0x1100 // ???
	case 0x3132:
		return 0x1101 // ???
	case 0x3134:
		return 0x1102 // ???
	case 0x3137:
		return 0x1103 // ???
	case 0x3138:
		return 0x1104 // ???
	case 0x3139:
		return 0x1105 // ???
	case 0x3141:
		return 0x1106 // ???
	case 0x3142:
		return 0x1107 // ???
	case 0x3143:
		return 0x1108 // ???
	case 0x3145:
		return 0x1109 // ???
	case 0x3146:
		return 0x110A // ???
	case 0x3147:
		return 0x110B // ???
	case 0x3148:
		return 0x110C // ???
	case 0x3149:
		return 0x110D // ???
	case 0x314A:
		return 0x110E // ???
	case 0x314B:
		return 0x110F // ???
	case 0x314C:
		return 0x1110 // ???
	case 0x314D:
		return 0x1111 // ???
	case 0x314E:
		return 0x1112 // ???
	}

	// The rest of the compatibility block consists of archaic compounds that are
	// unlikely to be encountered in modern systems. Just leave them alone.
	return jamo
}

// Transliterates one jamo at a time.
// Does nothing if it isn't in the modern jamo block.
func translitSingleJamo(jamo rune) string {
	jamo = decompatJamo(jamo)

	switch jamo {
	// Choseong (leading position consonants)
	case 0x1100:
		return "g" // ???
	case 0x1101:
		return "kk" // ???
	case 0x1102:
		return "n" // ???
	case 0x1103:
		return "d" // ???
	case 0x1104:
		return "tt" // ???
	case 0x1105:
		return "r" // ???
	case 0x1106:
		return "m" // ???
	case 0x1107:
		return "b" // ???
	case 0x1108:
		return "pp" // ???
	case 0x1109:
		return "s" // ???
	case 0x110A:
		return "ss" // ???
	case 0x110B:
		return "" // ???
	case 0x110C:
		return "j" // ???
	case 0x110D:
		return "jj" // ???
	case 0x110E:
		return "ch" // ???
	case 0x110F:
		return "k" // ???
	case 0x1110:
		return "t" // ???
	case 0x1111:
		return "p" // ???
	case 0x1112:
		return "h" // ???
	// Jungseong (vowels)
	case 0x1161:
		return "a" // ???
	case 0x1162:
		return "ae" // ???
	case 0x1163:
		return "ya" // ???
	case 0x1164:
		return "yae" // ???
	case 0x1165:
		return "eo" // ???
	case 0x1166:
		return "e" // ???
	case 0x1167:
		return "yeo" // ???
	case 0x1168:
		return "ye" // ???
	case 0x1169:
		return "o" // ???
	case 0x116A:
		return "wa" // ???
	case 0x116B:
		return "wae" // ???
	case 0x116C:
		return "oe" // ???
	case 0x116D:
		return "yo" // ???
	case 0x116E:
		return "u" // ???
	case 0x116F:
		return "wo" // ???
	case 0x1170:
		return "we" // ???
	case 0x1171:
		return "wi" // ???
	case 0x1172:
		return "yu" // ???
	case 0x1173:
		return "eu" // ???
	case 0x1174:
		return "ui" // ???
	case 0x1175:
		return "i" // ???
	// Jongseong (final position consonants)
	case 0x11A8:
		return "k" // ???
	case 0x11A9:
		return "k" // ???
	case 0x11AB:
		return "n" // ???
	case 0x11AE:
		return "t" // ???
	case 0x11AF:
		return "l" // ???
	case 0x11B7:
		return "m" // ???
	case 0x11B8:
		return "p" // ???
	case 0x11BA:
		return "t" // ???
	case 0x11BB:
		return "t" // ???
	case 0x11BC:
		return "ng" // ???
	case 0x11BD:
		return "t" // ???
	case 0x11BE:
		return "t" // ???
	case 0x11BF:
		return "k" // ???
	case 0x11C0:
		return "t" // ???
	case 0x11C1:
		return "p" // ???
	case 0x11C2:
		return "t" // ???
	}

	return string(jamo)
}

// Some combinations of ending jamo in one syllable and initial jamo in the next are romanized
// irregularly. These exceptions are called "special provisions". In cases where multiple
// romanizations are permitted, we use the one that's least commonly used elsewhere.
//
// Returns empty strring and false if either character is not in the modern jamo block,
// or if there is no special provision for that pair of jamo.
func translitSpecialProvisions(previousEnding rune, nextInitial rune) (string, bool) {
	// Return false if previousEnding not in modern jamo block
	if !unicode.In(previousEnding, jamoBlock) {
		return "", false
	}
	// Return false if nextInitial not in modern jamo block
	if !unicode.In(nextInitial, jamoBlock) {
		return "", false
	}

	// Jongseong (final position) ??? has a number of special provisions.
	if previousEnding == 0x11C2 {
		switch nextInitial {
		case 0x110B:
			return "h", true // ???
		case 0x1100:
			return "k", true // ???
		case 0x1102:
			return "nn", true // ???
		case 0x1103:
			return "t", true // ???
		case 0x1105:
			return "nn", true // ???
		case 0x1106:
			return "nm", true // ???
		case 0x1107:
			return "p", true // ???
		case 0x1109:
			return "hs", true // ???
		case 0x110C:
			return "ch", true // ???
		case 0x1112:
			return "t", true // ???
		default:
			return "", false
		}
	}

	// Otherwise, special provisions are denser when grouped by the second jamo.
	switch nextInitial {
	case 0x1100: // ???
		switch previousEnding {
		case 0x11AB:
			return "n-g", true // ???
		default:
			return "", false
		}
	case 0x1102: // ???
		switch previousEnding {
		case 0x11A8:
			return "ngn", true // ???
		case 0x11AE:
			fallthrough // ???
		case 0x11BA:
			fallthrough // ???
		case 0x11BD:
			fallthrough // ???
		case 0x11BE:
			fallthrough // ???
		case 0x11C0: // ???
			return "nn", true
		case 0x11AF:
			return "ll", true // ???
		case 0x11B8:
			return "mn", true // ???
		default:
			return "", false
		}
	case 0x1105: // ???
		switch previousEnding {
		case 0x11A8:
			fallthrough // ???
		case 0x11AB:
			fallthrough // ???
		case 0x11AF: // ???
			return "ll", true
		case 0x11AE:
			fallthrough // ???
		case 0x11BA:
			fallthrough // ???
		case 0x11BD:
			fallthrough // ???
		case 0x11BE:
			fallthrough // ???
		case 0x11C0: // ???
			return "nn", true
		case 0x11B7:
			fallthrough // ???
		case 0x11B8: // ???
			return "mn", true
		case 0x11BC:
			return "ngn", true // ???
		default:
			return "", false
		}
	case 0x1106: // ???
		switch previousEnding {
		case 0x11A8:
			return "ngm", true // ???
		case 0x11AE:
			fallthrough // ???
		case 0x11BA:
			fallthrough // ???
		case 0x11BD:
			fallthrough // ???
		case 0x11BE:
			fallthrough // ???
		case 0x11C0: // ???
			return "nm", true
		case 0x11B8:
			return "mm", true // ???
		default:
			return "", false
		}
	case 0x110B: // ???
		switch previousEnding {
		case 0x11A8:
			return "g", true // ???
		case 0x11AE:
			return "d", true // ???
		case 0x11AF:
			return "r", true // ???
		case 0x11B8:
			return "b", true // ???
		case 0x11BA:
			return "s", true // ???
		case 0x11BC:
			return "ng-", true // ???
		case 0x11BD:
			return "j", true // ???
		case 0x11BE:
			return "ch", true // ???
		default:
			return "", false
		}
	case 0x110F: // ???
		switch previousEnding {
		case 0x11A8:
			return "k-k", true // ???
		default:
			return "", false
		}
	case 0x1110: // ???
		switch previousEnding {
		case 0x11AE:
			fallthrough // ???
		case 0x11BA:
			fallthrough // ???
		case 0x11BD:
			fallthrough // ???
		case 0x11BE:
			fallthrough // ???
		case 0x11C0: // ???
			return "t-t", true
		default:
			return "", false
		}
	case 0x1111: // ???
		switch previousEnding {
		case 0x11B8:
			return "p-p", true // ???
		default:
			return "", false
		}
	default:
		return "", false
	}
}

// Decompose a syllable into several jamo. Does nothing if that isn't possible.
func decompose(syllable rune) string {
	return norm.NFD.String(string(syllable))
}

// Transliterate any Hangul in the given string.
// Leaves any non-Hangul characters unmodified.
func (kt *KoreanTranslit) Transliterate(s string) string {
	if len(s) == 0 {
		return s
	}

	builder := &strings.Builder{}

	nextInitialJamoConsumed := false

	for i, syllable := range s {
		// If character not in blocks, leave it unmodified
		if !unicode.In(syllable, jamoBlock, syllablesBlock, compatJamoBlock) {
			builder.WriteRune(syllable)
			continue
		}

		jamo := decompose(syllable)
		for j, char := range jamo {
			// If we already transliterated the first jamo of this syllable as part of a special
			// provision, skip it. Otherwise, handle it in the unconditional else branch.
			if j == 0 && nextInitialJamoConsumed {
				nextInitialJamoConsumed = false
				continue
			}

			// If this is the last jamo of this syllable and not the last syllable of the
			// string, check for special provisions. If the next char is whitespace or not
			// Hangul, run translitSpecialProvisions() should return no value.
			if j == len(jamo)-1 && i < len(s)-1 {
				nextSyllable := s[i+1]
				nextJamo := decompose(rune(nextSyllable))[0]

				// Attempt to handle special provision
				specialProvision, ok := translitSpecialProvisions(char, rune(nextJamo))
				if ok {
					builder.WriteString(specialProvision)
					nextInitialJamoConsumed = true
				} else {
					// Not a special provision, transliterate normally
					builder.WriteString(translitSingleJamo(char))
				}
				continue
			}
			// Transliterate normally
			builder.WriteString(translitSingleJamo(char))
		}
	}
	return builder.String()
}
