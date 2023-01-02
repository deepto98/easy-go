package main

import (
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
		Debug:   true,
	}
	err = easyGoApp.New(path)

	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		App: easyGoApp,
	}
	return app
}
