/*

  Mayor elemento: dado un arreglo de enteros, devolver el mayor elemento

*/

package tp2

func greaterElement(arr []int, gt int) int {
  if len(arr) == 0 {
    return gt
  }

  if gt < arr[0] {
    gt = arr[0]
  }

  return greaterElement(arr[1:], gt)
}

func GreaterElement(arr []int) int {
  return greaterElement(arr, 0)
}
