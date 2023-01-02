package easygo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cuelang.org/go/pkg/strconv"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

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

	//Creates folders and files
	err := easyGo.Init(pathConfig)

	if err != nil {
		return err
	}

	//Adds logging
	easyGo.InfoLog, easyGo.ErrorLog = easyGo.InitLoggers()

	easyGo.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG_MODE"))

	easyGo.Version = version

	easyGo.RootPath = rootPath

	easyGo.Routes = easyGo.routes().(*chi.Mux)

	easyGo.config = Config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
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

//Initializes custom error and info loggers
func (easyGo *EasyGo) InitLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

//Starts the web server
func (easyGo *EasyGo) ListenAndServe() {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", easyGo.config.port),
		ErrorLog:     easyGo.ErrorLog,
		Handler:      easyGo.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	easyGo.InfoLog.Println("Listening on port :", easyGo.config.port)

	err := srv.ListenAndServe()
	easyGo.ErrorLog.Fatal(err)

}
