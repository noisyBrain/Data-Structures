package main

import "fmt"

/*
  Calcule el tiempo de ejecución para los siguientes algoritmos suponiendo que cada
  operación elemental tarda un tiempo constante. Luego indique el Orden en el peor
  caso para cada uno de ellos.
*/

func ejemplo1(n int) {
  for i := 0; i < n; i++ {
    fmt.Println(i)
  }
} // lineal -> O(n)

func ejemplo2(n int, m int) {
  for i := 0; i < n; i++ {
    fmt.Printf("(i: %v)", i)
  }
  for j := 0; j < m; j++ {
    fmt.Printf("(j: %v)", j)
  }
} // lineal -> O(n + m)
  // simplificando queda en O(n)

func ejemplo3(n int) {
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      fmt.Printf("(i: %v, j: %v)", i, j)
    }
  }
} // cuadrática O(n²)

func ejemplo4(n int) {
  for i := 0; i < n; i++ {
    for j := i; j < n; j++ {
      fmt.Printf("(i: %v, j: %v)", i, j)
    }
  }
} // cuadrática porque:
  // [n.(n - 1)] / 2 = (n² - n) / 2
  // simplificando queda de O(n2)