package main

import (
	"github.com/ka1i/InnovationCenter/internal/graphical"
)

func main() {
	flags()

	app := graphical.UserInterface()
	app.Run()
}
