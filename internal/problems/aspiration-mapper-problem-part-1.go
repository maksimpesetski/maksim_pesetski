package problems

import (
	"unicode"
)

/*
 Problem one:
	Write a function that accepts a sting input and capitalizes *only* every third alphanumeric character of a string,
	so that if I hand you "Aspiration.com"
	You hand me back "asPirAtiOn.cOm"
	Please note:
	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].
*/

// CapitalizeEveryThirdAlphanumericChar accepts a sting input and returns a new string where *only* every third
// alphanumeric character of a string is capitalized. Time: O(n), Space: O(n) where n - number chars in the string.
func CapitalizeEveryThirdAlphanumericChar(s string) string {

	// convert sting value into rune in order to get access to methods like .IsLetter().
	result := []rune(s)

	// initialize the letter count to keep track of visited alphabetical characters count.
	letterCount := 0

	// loop over every rune
	for i, r := range result {

		// check if the current rune is letter
		if unicode.IsLetter(r) {

			// increment letter counter
			letterCount++

			// capitalize current rune if it's an every third alphanumeric character
			if letterCount%3 == 0 {
				result[i] = unicode.ToUpper(r)

				// downcase the current rune if it's NOT an every third alphanumeric character
			} else {
				result[i] = unicode.ToLower(r)
			}

		}
	}

	// converted slice of runes into sting and return it
	return string(result)
}
