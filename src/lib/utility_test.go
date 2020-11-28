package lib

import (
	"errors"
	"practice/demo/src/entity"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	t.Run("json file reading error", func(t *testing.T) {
		old := readFile
		readFile = func(filename string) ([]byte, error) {
			return nil, errors.New("error in reading file")
		}
		defer func() {
			readFile = old
		}()
		err := ReadJSONFile("abc.json", nil)
		if err == nil {
			t.Fatal("json file read should return error")
		}
	})
	t.Run("json file read success", func(t *testing.T) {
		filePath := "../../config.json"
		obj := &entity.AppConfig{}
		err := ReadJSONFile(filePath, obj)
		if err != nil {
			t.Fatalf("json file read should return error succcess = %v", err)
		}
		if len(obj.ListenURL) < 1 {
			t.Fatal("error in reading file")
		}
	})
}
