package service

import (
	"os"
)

func GetCollections() (collections []string) {
	files, _ := os.ReadDir("/collections")
	for _, file := range files {
		collections = append(collections, file.Name())
	}
	return
}

func CreateCollection(name string) error {
	_, err := os.Create("/collections/" + name)
	return err
	//return os.WriteFile("/collections/"+name, nil, 0666)
}

func DeleteCollection(name string) error {
	if _, err := os.Open("/collections/" + name); err != nil {
		// Handle error
	}

	if err := os.Remove("/collections/" + name); err != nil {
		// Handle error
	}
	return nil
}

func GetCollection(name string) {

}
