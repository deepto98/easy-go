package main

import easygo "github.com/deepto98/easy-go"

type Application struct {
	App *easygo.EasyGo
}

func main() {
	easyGo := initApp()
	easyGo.App.ListenAndServe()
}
