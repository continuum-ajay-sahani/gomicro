package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
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

// ParseRequestBody parse request query parameter from request body(POST, PUT, DELETE)
func ParseRequestBody(r *http.Request, m interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	bv := make(map[string]string)
	err = json.Unmarshal(body, &bv)
	if err != nil {
		panic(err)
	}
	q, _ := bv["query"]
	json.Unmarshal([]byte(q), &m)
}
