package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Molsbee/mock-server/model"
)

type CollectionRepo interface {
	GetCollectionNames() []string
	GetCollections() []model.Collection
	GetCollection(string) (*model.Collection, error)
	CreateCollection(model.Collection) error
	UpdateCollection(string, model.Collection) (*model.Collection, error)
	DeleteCollection(string) error
}

type fileCollectionRepo struct {
	collectionsDirectory string
}

func NewFileCollectionRepo() CollectionRepo {
	currentDirectory, _ := os.Getwd()
	collectionsDirectory := currentDirectory + "/collections"

	if _, err := os.Stat(collectionsDirectory); err != nil {
		if dirErr := os.Mkdir(collectionsDirectory, 0777); dirErr != nil {
			log.Panicln("failed to initialize the collections directory")
		}
	}

	return &fileCollectionRepo{
		collectionsDirectory: collectionsDirectory,
	}
}

func (fcr *fileCollectionRepo) GetCollectionNames() []string {
	files, _ := os.ReadDir(fcr.collectionsDirectory)
	collections := make([]string, len(files))
	for i, file := range files {
		parts := strings.Split(file.Name(), ".json")
		collections[i] = parts[0]
	}
	return collections
}

func (fcr *fileCollectionRepo) GetCollections() []model.Collection {
	files, _ := os.ReadDir(fcr.collectionsDirectory)
	collections := make([]model.Collection, len(files))
	for i, file := range files {
		collection := model.Collection{Name: file.Name()}
		data, err := os.ReadFile(fcr.collectionsDirectory + "/" + file.Name())
		if err != nil {
			log.Printf("failed to read %s collection file - %s", file.Name(), err.Error())
			continue
		}
		if err := json.Unmarshal(data, &collection); err != nil {
			log.Printf("failed to unmarshal %s collection data - %s", file.Name(), err.Error())
			continue
		}
		collections[i] = collection
	}
	return collections
}

func (fcr *fileCollectionRepo) CreateCollection(collection model.Collection) error {
	_, err := os.Stat(fmt.Sprintf("%s/%s.json", fcr.collectionsDirectory, collection.Name))
	if err == nil {
		return fmt.Errorf("collection %s already exists", collection.Name)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.json", fcr.collectionsDirectory, collection.Name))
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

func (fcr *fileCollectionRepo) UpdateCollection(name string, collection model.Collection) (*model.Collection, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s.json", fcr.collectionsDirectory, name))
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

func (fcr *fileCollectionRepo) DeleteCollection(name string) error {
	_, err := os.Stat(fcr.collectionsDirectory + "/" + name)
	if err == nil {
		return os.Remove(fcr.collectionsDirectory + "/" + name)
	}
	return nil
}

func (fcr *fileCollectionRepo) GetCollection(name string) (*model.Collection, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s.json", fcr.collectionsDirectory, name))
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
