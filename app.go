package easygo

const (
	version = "1.0.0"
)

type EasyGo struct {
	AppName string
	Debug   bool
	Version string
}

//Creates new project
func (easyGo *EasyGo) New(rootPath string) error {

}

//Initializes folder structure for go project
func (easyGo *EasyGo) Init(paths InitPaths) error {
	rootPath := paths.rootPath

	for _, path := range paths.folderNames {
		//Todo : Add func to create folder
	}
}
