/*

  Posición mayúscula: dada una cadena de caracteres, devuelve la posición en la que se
  encuentra la primera letra mayúscula.

*/

package tp2

import "unicode"

func FirstMayusPosition(str string) int {
  if len(str) == 0 {
    return -1
  }

  isUpper := unicode.IsUpper(rune(str[0]))
  if isUpper {
    return 0
  }

  index := FirstMayusPosition(str[1:])

  if index == -1 {
    return -1
  }

  return index + 1
} 
