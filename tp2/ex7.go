/*

  Invertir: dado un arreglo de enteros, invertirlo.

*/

package tp2

func invert(arr []int, aux []int) []int {
  if len(arr) == 0 {
    return aux
  }

  aux = append(aux, arr[len(arr) - 1])
  
  return invert(arr[:len(arr) - 1], aux);
}

func Invert(arr []int) []int {
  array := []int {}
  return invert(arr, array)
}
