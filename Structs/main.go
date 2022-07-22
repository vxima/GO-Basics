package main

import "fmt"

func main() {
	type Jogo struct {
		nome  string
		nota  float64
		steam bool
	}
	elden := Jogo{nome: "Elden Ring", nota: 10.0, steam: true}
	fmt.Println(elden.nome)
}
