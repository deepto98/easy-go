package main

import (
	"demo/handlers"

	easygo "github.com/deepto98/easy-go"
)

type Application struct {
	App      *easygo.EasyGo
	Handlers *handlers.Handlers
}

func main() {
	easyGo := initApp()
	easyGo.App.ListenAndServe()
}
