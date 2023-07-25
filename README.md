# Saludos en Go

Este es un programa simple que imprime un saludo en Go.

## Instalación

Para instalarlo, usa `go get`:

```bash
go get -u github.com/jotafierro/greetings
```

## Ejemplo de uso

```go
package main

import (
  "fmt"

  "github.com/jotafierro/greetings"
)

func main() {
  // Obtén un saludo aleatorio en inglés
  message, err := greetings.Hello("Jonathan Fierro")
  if err != nil {
    log.Fatal(err)
  }

  // Imprime el saludo en la consola
  fmt.Println(message)
}
```