package easygo

import "github.com/joho/godotenv"

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

//Initializes folder structure and creates required files for go project
func (easyGo *EasyGo) Init(paths InitPaths) error {
	root := paths.rootPath

	// 1. Create necessary folders
	for _, path := range paths.folderNames {
		err := easyGo.CreateDirectoryIfDoesNotExist(root + "/" + path)

		if err != nil {
			return err
		}
	}

	// 2. Check for environment file, create if it doesn't exit
	err := easyGo.CheckEnvFile(root)

	if err != nil {
		return err
	}
	// 3. Read Env file
	err = godotenv.Load(root + "/.env")

	if err != nil {
		return err
	}

	return nil
}

func (easyGo *EasyGo) CheckEnvFile(path string) error {
	err := easyGo.CreateFileIfDoesNotExist(path + "/.env")

	if err != nil {
		return err
	}
	return nil
}
