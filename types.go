package easygo

import "log"

//Root type for App
type EasyGo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	config   Config
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
