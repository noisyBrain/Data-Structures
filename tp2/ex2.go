/*

  Suma de enteros: permite sumar todos los números enteros comprendidos entre un
  parámetro de inicio y uno de fin.

*/

package tp2;

func SumInts(num1 int, num2 int) int {
  if num1 == num2 || num1 > num2 {
    return 0;
  }

  return num1 + SumInts(num1+1, num2);
}
