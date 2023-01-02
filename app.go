package easygo

const (
	version = "1.0.0"
)

var folderNames = []string{
	"handlers", "migrations", "views", "data", "public", "logs", "middleware",
}

//Creates new Go project
func (easyGo *EasyGo) New(rootPath string) error {
	pathConfig := InitPaths{
		rootPath:    rootPath,
		folderNames: folderNames,
	}
	err := easyGo.Init(pathConfig)

	if err != nil {
		return err
	}
	return nil
}

//Initializes folder structure for go project
func (easyGo *EasyGo) Init(paths InitPaths) error {
	root := paths.rootPath

	for _, path := range paths.folderNames {
		err := easyGo.CreateDirectoryIfDoesNotExist(root + "/" + path)

		if err != nil {
			return err
		}
	}
	return nil
}
