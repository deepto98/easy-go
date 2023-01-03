package main

import (
	"demo/handlers"
	"log"
	"os"

	easygo "github.com/deepto98/easy-go"
)

func initApp() *Application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	easyGoApp := &easygo.EasyGo{
		AppName: "demo-app",
		// Debug:   true,
	}
	err = easyGoApp.New(path)

	if err != nil {
		log.Fatal(err)
	}

	handlers := &handlers.Handlers{
		App: easyGoApp,
	}

	app := &Application{
		App:      easyGoApp,
		Handlers: handlers,
	}
	app.App.Routes = app.routes()
	easyGoApp.InfoLog.Println("Debug:", easyGoApp.Debug)

	return app
}
