package persist

import (
	"fmt"
	"os"
	pathFs "path"
	"strings"
)

const INTERNAL_PATH = ".plaintext-db"
var DATABASES_INFO_PATH = pathFs.Join(INTERNAL_PATH, "databases.dat")

type PlainTextDb struct {
	path string
}

func NewPlainTextDb(path string) (*PlainTextDb, error) {
	if _, err := os.Stat(INTERNAL_PATH); os.IsNotExist(err) {
		err = os.MkdirAll(INTERNAL_PATH, 0755)
		if err != nil {
			return nil, err
		}
	}

	err := appendOrCreate(DATABASES_INFO_PATH, path)
	if err != nil {
		return nil, err
	}

	return &PlainTextDb{
		path: path,
	}, nil
}

func (db *PlainTextDb) Save(data string) (error) {
	fullpath := pathFs.Join(INTERNAL_PATH, db.path)
	return appendOrCreate(fullpath, data)
}

func (db *PlainTextDb) WipeSingle(path string) (error) {
	fullpath := pathFs.Join(INTERNAL_PATH, path)
	if _, err := os.Stat(fullpath); os.IsExist(err) {
		err = os.Remove(fullpath)
		return err
	}

	return nil
}

func (db *PlainTextDb) Wipe() (error) {
	data, err := os.ReadFile(DATABASES_INFO_PATH)
	if err != nil {
		return err
	}

	for _, database := range strings.Split(string(data), "\n") {
		err = db.WipeSingle(database)
		if err != nil {
			return err
		}
	}

	return nil
}

func appendOrCreate(path string, data string) (error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data = data + "\n"
	_, err = file.WriteString(data)

	return err
}