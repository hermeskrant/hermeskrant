package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	infos, err := ioutil.ReadDir("_data/articles")
	if err != nil {
		fmt.Println(err.Error())
	}

	var files = make([]interface{}, len(infos))

	for i, info := range infos {
		file, err := ioutil.ReadFile(filepath.Join("_data/articles", info.Name()))
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.Unmarshal(file, &files[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		files[i].(map[string]interface{})["path"] = info.Name()
	}

	output, err := json.Marshal(files)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile("_data/articles.json", output, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
