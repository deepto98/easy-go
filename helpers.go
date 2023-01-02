package easygo

import "os"

func (easyGo *EasyGo) CreateDirectoryIfDoesNotExist(path string) error {
	const permissions = 0755

	//Check if folder doesn't exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, permissions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (easyGo *EasyGo) CreateFileIfDoesNotExist(path string) error {

	//Check if file doesn't exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)

		if err != nil {
			return err
		}

		defer file.Close()
	}
	return nil
}
