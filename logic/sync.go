package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//Copy file from localPath to remotePath
	localPath := "C:/GoStudio2"
	remotePath := "C:/GoStudio3"

	err := syncFiles(localPath, remotePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File synchronization completed successfully")
}

func syncFiles(localPath, remotePath string) error {
	files, err := ioutil.ReadDir(localPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		localFilePath := filepath.Join(localPath, file.Name())
		remoteFilePath := filepath.Join(remotePath, file.Name())

		if file.IsDir() {
			err = os.MkdirAll(remoteFilePath, os.ModePerm)
			if err != nil {
				return err
			}

			err = syncFiles(localFilePath, remoteFilePath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(localFilePath, remoteFilePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
