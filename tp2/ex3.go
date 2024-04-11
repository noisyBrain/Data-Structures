/*

  Factorial de un número entero.

*/

package tp2

func Factorial(num int) int {
  if num == 0 {
    return 0;
  }

  if num == 1 {
    return 1;
  }

  return num * Factorial(num - 1);
} 
