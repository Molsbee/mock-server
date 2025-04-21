package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Molsbee/mock-server/model"
)

var collectionsDirectory string

func init() {
	currentDirectory, _ := os.Getwd()
	collectionsDirectory = currentDirectory + "/collections"

	if _, err := os.Stat(collectionsDirectory); err != nil {
		if dirErr := os.Mkdir(collectionsDirectory, 0777); dirErr != nil {
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

func CreateCollection(collection model.Collection) error {
	file, err := os.Create(fmt.Sprintf("%s/%s.json", collectionsDirectory, collection.Name))
	if err != nil {
		return fmt.Errorf("failed to create %s collection file - (%s)", collection.Name, err.Error())
	}
	defer file.Close()

	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal %s collection data - (%s)", collection.Name, err.Error())
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write %s collection data - (%s)", collection.Name, err.Error())
	}

	return nil
}

func UpdateCollection(name string, collection model.Collection) (*model.Collection, error) {
	file, err := os.Open(collectionsDirectory + "/" + name)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s collection file - %s", name, err.Error())
	}
	defer file.Close()

	data, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s collection data - %s", collection.Name, err.Error())
	}

	_, err = file.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write %s collection data - %s", collection.Name, err.Error())
	}
	return &collection, nil
}

func DeleteCollection(name string) error {
	_, err := os.Stat(collectionsDirectory + "/" + name)
	if err == nil {
		return os.Remove(collectionsDirectory + "/" + name)
	}
	return nil
}

func GetCollection(name string) (*model.Collection, error) {
	file, err := os.Open(collectionsDirectory + "/" + name)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s collection file - %s", name, err.Error())
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s collection file - %s", name, err.Error())
	}

	var collection model.Collection
	if err := json.Unmarshal(data, &collection); err != nil {
		return nil, fmt.Errorf("failed to convert %s content to collection struct - %s", name, err.Error())
	}
	return &collection, nil
}
