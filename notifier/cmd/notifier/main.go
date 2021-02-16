package main

import (
	"github.com/ivyoverflow/pub-sub/notifier/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
