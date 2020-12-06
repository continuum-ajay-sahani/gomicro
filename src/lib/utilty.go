package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gocql/gocql"
)

// TransactionContext use data context
type TransactionContext struct{}

// TransactionContextKey used data context key
var TransactionContextKey = TransactionContext{}

// ReadJSONFile reads config data from JSON-file
var ReadJSONFile = func(filePath string, obj interface{}) error {
	log.Printf("looking for JSON file to read (%s)", filePath)

	if len(filePath) < 1 {
		return errors.New("please specify file path to read")
	}
	contents, err := ReadFile(filePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(obj)
	}
	if err != nil {
		log.Printf("reading JSON file (%s) failed: %s\n", filePath, err)
	} else {
		log.Printf("read from JSON (%s) successfully\n", filePath)
	}

	return err
}

// ReadFile reads file from the given path
var ReadFile = func(filename string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Clean(filename))
}

// GetTransactionContextValue get context value
func GetTransactionContextValue(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	v := ctx.Value(TransactionContextKey)
	if v == nil {
		return ""
	}

	return v.(string)
}

// GetUniqueID return unique id
func GetUniqueID() string {
	return gocql.TimeUUID().String()
}
