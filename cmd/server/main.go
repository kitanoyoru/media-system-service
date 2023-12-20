package main

import (
	"github.com/kitanoyoru/media-system-service/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
