# Onda Vital

El propsito es hacer un crawl a wikipedia, y buscar en la barra de informacion el titulo de peliculas en España.
Si a primera instancia no encuentra la pagina, sigue los links de recomendaciones en go routines y tambien busca en ellos.

## Rant
Esta es una libreria en GO, principalmente para familiarizarme con el lenguaje.
el codigo es horrible, intente demasiadas idea y experimente bastante con este.
Aun asi subo para probar hacer un package + import a github.

## Rant++
En serio, hay hacks sobre hacks, principalmente al buscar links, y con los channels para sincronizar el crawling de sugerencias.


### Ejemplo:
```Go
package main

import (
  "github.com/ozkar99/ondavital"
  "fmt"
  )

func main() {
  val, err := ondavital.Search("die hard")
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(val)
}

//> Jugla de Cristal
// WTF?
```


![Img](http://i.imgur.com/Zt0T4SY.png)
