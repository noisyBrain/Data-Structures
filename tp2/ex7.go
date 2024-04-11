/*

  Invertir: dado un arreglo de enteros, invertirlo.

*/

package tp2

func invertir(arr []int, aux []int) []int {
  if len(arr) == 0 {
    return aux
  }

  aux = append(aux, arr[len(arr) - 1])
  
  return invertir(arr[:len(arr) - 1], aux);
}

func Invertir(arr []int) []int {
  array := []int {}
  return invertir(arr, array)
}
