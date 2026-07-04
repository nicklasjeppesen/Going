package container

import (
	"fmt"

	helper "myapp/internal/app/helper"

	customrouter "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

type Logger struct {
}

func (l *Logger) Log(message string) {
	fmt.Println(message)
}

func GetContainer() *customrouter.Container {

	container := customrouter.NewContainer()
	container.Register((*helper.ILogger)(nil), Logger{})
	return container

}
