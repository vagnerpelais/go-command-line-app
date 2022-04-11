package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.AddressSearch()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
