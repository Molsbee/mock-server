package service

import (
	"log"
	"os"
)

var collectionsDirectory string

func init() {
	currentDirectory, _ := os.Getwd()
	collectionsDirectory = currentDirectory + "/collections"

	if _, err := os.Stat(collectionsDirectory); err != nil {
		if dirErr := os.Mkdir(collectionsDirectory, 0666); dirErr != nil {
			log.Panicln("failed to initialize the collections directory")
		}
	}
}

func GetCollections() []string {
	files, _ := os.ReadDir(collectionsDirectory)
	collections := make([]string, len(files))
	for i, file := range files {
		collections[i] = file.Name()
	}
	return collections
}

func CreateCollection(name string) error {
	err := os.Mkdir(collectionsDirectory+"/"+name, 0666)
	return err
}

func DeleteCollection(name string) error {
	_, err := os.Stat(collectionsDirectory + "/" + name)
	if err == nil {
		return os.Remove(collectionsDirectory + "/" + name)
	}
	return nil
}

func GetCollection(name string) {

}
