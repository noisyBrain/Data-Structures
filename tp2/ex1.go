/*

  Cuenta regresiva: recibe un número entero e imprime todos los números comprendidos
  entre el mismo y 0.

*/

package main;

import "fmt";

func countDown(num int) int {
  if num < 0 {
    return 0;
  };

  fmt.Println(num);
  return countDown(num - 1);
}

func takeInput(input *int) *int  {
  fmt.Println("Ingresá un número: ");
  _, err := fmt.Scanln(input);

  if err != nil {
    fmt.Println("Error al leer la entrada", err);
  }

  return input;
}

func main()  {
  var input int;
  takeInput(&input);

  countDown(input);
}
