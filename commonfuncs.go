package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func arrayContains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ifs(condition bool, val1, val2 string) string {
	if condition {
		return val1
	}
	return val2
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func readJSONfile(filename string, data interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &data)
	return err
}

func curDir() string {
	dir, _ := os.Getwd()
	return dir
}