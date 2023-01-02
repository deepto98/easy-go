package easygo

//Root type for App
type EasyGo struct {
	AppName string
	Debug   bool
	Version string
}

//Type to create config with root directory and subfolder names which will
//be created during initialization
type InitPaths struct {
	rootPath    string
	folderNames []string
}
