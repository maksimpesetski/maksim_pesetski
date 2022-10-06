package problems

import (
	"unicode"
)

/*
	Problem 2:
		Do the same problem, but this time create a "mapper" package that looks like this:
		package mapper

		type Interface interface {
			TransformRune(pos int)
			GetValueAsRuneSlice() []rune
		}

		func MapString(i Interface) {
			for pos := range i.GetValueAsRuneSlice() {
				i.TransformRune(pos)
			}
		}
*/

// SkipSting is wrapper around a string type that has custom methods in it
type SkipSting struct {
	value             []rune
	step              int   // identifies the step for "every" char we want to capitalize e.g. 3.
	cashedLetterCount []int // serves as a caching layer to reduce duplicated alphabetical chars count calculation
}

/*
	NOTE:
		We could refactor NewCapitalizeEveryThirdAlphanumericCharSting to return mapper.Mapper type in order to make
		"SkipSting" a private pkg type and only expose external mapper package. The current pkg structure keeps packages
		isolated.
*/
// NewSkipString accepts step int and string inputs and creates a new SkipSting. Time: O(1); Space: O(n);
func NewSkipString(step int, s string) SkipSting {
	return SkipSting{
		value:             []rune(s),
		step:              step,
		cashedLetterCount: make([]int, len(s)),
	}
}

// GetValueAsRuneSlice returns the SkipSting value as a []rune. Time: O(1); Space: O(1);
func (p SkipSting) GetValueAsRuneSlice() []rune {
	return p.value
}

/*
	NOTE:
		We could simply call CapitalizeEveryThirdAlphanumericChar from aspiration-mapper-problem-part-1.go to reduce
		code duplication, since it's covered with unit-tests, but this would increase space complexity.
*/
// TransformRune mutates the original SkipSting at passed pos value by capitalizing *only* every step alphanumeric character.
// Time: O(n); Space: O(1) where n - number of characters in sting.
func (p *SkipSting) TransformRune(pos int) {

	// edge case if pos is outbound
	if pos < 0 || pos >= len(p.value) {
		return
	}

	var lettersCount int

	// iterate over the value up to passed position to calculate alphabetical char instances
	for i := 0; i <= pos; i++ {

		// check if we already calculated the alphabetic characters count before by looking up the cache
		if p.cashedLetterCount[i] != 0 {
			// grab cached alphabetic characters count
			lettersCount = p.cashedLetterCount[i]
			continue
		}

		// cache miss

		// check if the char is alphabetic
		if unicode.IsLetter(p.value[i]) {

			// increment alphabetic characters count for uncached value
			lettersCount++

			// cache new alphabetic characters count for this value
			p.cashedLetterCount[i] = lettersCount

			// only modify the value for the passed position
			if i == pos {
				// modify the original SkipSting value based on the following conditions
				//capitalize current rune if it's an every "step" alphanumeric character
				if (lettersCount % p.step) == 0 {
					p.value[pos] = unicode.ToUpper(p.value[pos])

					// downcase the current rune if it's NOT an every third alphanumeric character
				} else {
					p.value[pos] = unicode.ToLower(p.value[pos])
				}
			}

			// if not alphabetic copy previously cashed value
		} else {
			if i-1 >= 0 {
				p.cashedLetterCount[i] = p.cashedLetterCount[i-1]
			}
		}
	}
}

// String implements the Stringer interface. Time: O(1); Space: O(1);
func (p SkipSting) String() string {
	return string(p.value)
}
