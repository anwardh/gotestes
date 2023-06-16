package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

const (
	FileType string = "arquivo"
)

type Mock struct {
	Data []byte // json
	Err  error
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

func (fs *FileStore) AddMock(mock *Mock) {
	fs.Mock = mock
}

func (fs *FileStore) ClearMock() {
	fs.Mock = nil
}

// Estamos criando a nossa strict de Mock

func Factory(store string, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{
			fileName,
			nil,
		}
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		return nil
	}

	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	// Esse trecho de código pergunta se a pessoa quis usar um mock
	if fs.Mock != nil {
		if fs.Mock.Err != nil {
			return fs.Mock.Err
		}
		// Caso tenha usado o mock, ele retorna o dado que a pessoa passou no mock
		return json.Unmarshal(fs.Mock.Data, data)
	}

	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}
