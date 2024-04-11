package main

import (
	"fmt"
	"tps/tp2"
)

func main() {
  // ex 2
  fmt.Println("\nExercise 2:")
  fmt.Println(tp2.SumInts(1, 3))

  // ex 3
  fmt.Println("\nExercise 3:")
  fmt.Println(tp2.Factorial(5))

  // ex 4
  fmt.Println("\nExercise 4:")
  fmt.Println(tp2.VowelsAndConsonants("#testeando <la> @ funcion!!")) // should return 8 vowels, 10 consonants

  // ex 5
  fmt.Println("\nExercise 5:")
  fmt.Println(tp2.FirstMayusPosition("hola como Te va")) // should return 10

  // ex 6
  fmt.Println("\nExercise 6:")
  arr := []int {0, 6, 2, 3}
  fmt.Println(tp2.GreaterElement(arr)) // should return 6
}
