package container

import (
	"log"
	helper "myapp/internal/app/helper"

	customrouter "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

type Logger struct {
}

func (l *Logger) Log(message string) {
	log.Println(message)
}

func GetContainer() *customrouter.Container {

	container := customrouter.NewContainer()

	container.Register(func() helper.ILogger {
		return new(Logger)
	})

	return container

}
