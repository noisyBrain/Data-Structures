package main

/* 
  Dado el siguiente algoritmo:

*/

  func buscar(lista []string, buscado string) int {
    for i := 0; i < len(lista); i++ {
      if lista[i] == buscado {
        return i
      }
    }

    return -1
  }

/*
  a) Calcule T(n) en el peor caso. ¿Cuál es el Orden del algoritmo? O(n)

  b) ¿Podemos decir también que el algoritmo tiene O(n²)? ¿Por qué? 
     No. Porque la complejidad aumenta a medida que se añaden elementos a la lista y en ningún momento se la va a recorrer dos veces

  c) ¿Podemos decir también que el algoritmo tiene O(log² n)? ¿Por qué?
     No. Para ser un log² n, la búsqueda debería dividirse a la mitad en cada iteración, así como lo hace Binary Search, donde en cada iteración se divide la cantidad de elementos a la mitad y se descartar una de ellas en base al número que se está buscando. 

  d) ¿Podemos decir también que el algoritmo tiene O(1)? ¿Por qué?
     No. Porque para ser constante debería ejecutarse un sólo paso sin importar la cantidad de elementos en la lista.
*/
