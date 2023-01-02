package easygo

import (
	"log"

	"github.com/deepto98/easy-go/render"
	"github.com/go-chi/chi/v5"
)

//Root type for App
type EasyGo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   Config
	Routes   *chi.Mux
	Render   *render.Render
}

//Type to create config with root directory and subfolder names which will
//be created during initialization
type InitPaths struct {
	rootPath    string
	folderNames []string
}

//Stores app config
type Config struct {
	port     string
	renderer string
}
