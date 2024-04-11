/*

  Vocales y consonantes: devuelve la cantidad de vocales y de consonantes que contiene una cadena.

*/

package tp2

import (
	"strings"
)

func VowelsAndConsonants(str string) (int, int) {
  if len(str) == 0 {
    return 0, 0
  }

  totalVowels, totalConsonants := VowelsAndConsonants(str[1:])
  vowels := "aeiou"
  specialChars := "!@#$%^&*()<>¿?"

  currentChar := string(str[0])
  isSpecialChar := currentChar == " " || strings.Contains(specialChars, currentChar)
  isVowel := strings.Contains(vowels, currentChar)

  if isSpecialChar {
    return totalVowels, totalConsonants
  }

  if isVowel {
    return totalVowels + 1, totalConsonants
  }

  return totalVowels, totalConsonants + 1
}
