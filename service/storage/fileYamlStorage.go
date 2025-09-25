package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

type FileYamlStorage[T any] struct {
	BasePath string
	File     string
	locker   sync.Mutex
}

func (f *FileYamlStorage[T]) getFilePath() string {
	return filepath.Join(f.BasePath, f.File+".yaml")

}
func NewFileYamlStorage[T any](basePath string, file string) *FileYamlStorage[T] {
	return &FileYamlStorage[T]{BasePath: basePath, File: file, locker: sync.Mutex{}}
}

// Load reads a YAML file identified by a key and unmarshals its content
// into the provided entity. The entity argument should be a pointer to the
// target data structure (e.g., &myStruct).
func (f *FileYamlStorage[T]) Load(entity *T) error {
	// Construct the full file path from the base path and key.
	filePath := f.getFilePath()

	// Read the entire file content.
	data, err := os.ReadFile(filePath)
	if err != nil {
		// Handle file not found as a specific case, otherwise return a generic read error.
		if os.IsNotExist(err) {
			return NotFoundErr
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the YAML data into the provided entity pointer.
	if err := yaml.Unmarshal(data, entity); err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return nil
}

// Save marshals the given entity into YAML format and writes it to a file
// identified by the key.
func (f *FileYamlStorage[T]) Save(entity T) error {
	// Marshal the entity into YAML format.
	data, err := yaml.Marshal(entity)
	if err != nil {
		return fmt.Errorf("failed to marshal entity to yaml: %w", err)
	}

	_ = os.MkdirAll(filepath.Dir(f.BasePath), 0755)

	// Construct the full file path.
	filePath := f.getFilePath()

	// Write the data to the file, creating it if it doesn't exist.
	// 0644 provides standard read/write permissions.
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (f *FileYamlStorage[T]) Set(entity T) error {
	return f.Save(entity)
}

func (f *FileYamlStorage[T]) Get() (T, error) {
	data := new(T)
	err := f.Load(data)
	return *data, err

}

func (f *FileYamlStorage[T]) Transaction(tx func(entity *T, loadError error) error) error {
	f.locker.Lock()
	defer f.locker.Unlock()

	entity := new(T)
	err := f.Load(entity)

	if err := tx(entity, err); err != nil {
		return err
	}

	if err := f.Save(*entity); err != nil {
		return err
	}

	return nil

}
