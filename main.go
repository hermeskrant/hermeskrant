package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

var months = []string{
	"januari",
	"februari",
	"maart",
	"april",
	"mei",
	"juni",
	"juli",
	"augustus",
	"september",
	"oktober",
	"november",
	"december",
}

func main() {
	infos, err := ioutil.ReadDir("_data/articles")
	if err != nil {
		fmt.Println(err.Error())
	}

	var files = make([]interface{}, len(infos))
	var final = make(map[string]interface{}, 1)

	for i, info := range infos {
		file, err := ioutil.ReadFile(filepath.Join("_data/articles", info.Name()))
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.Unmarshal(file, &files[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		files[i].(map[string]interface{})["path"] = strings.Replace(info.Name(), ".json", "", -1)
		if files[i].(map[string]interface{})["date"] == nil {
			t := time.Now()
			year, month, day := t.Date()
			files[i].(map[string]interface{})["date"] = fmt.Sprintf("%v %v %v", day, months[month], year)
		}

		outputJSON, err := json.Marshal(files[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		err = ioutil.WriteFile("_data/articles/"+info.Name(), outputJSON, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	final["Articles"] = files

	output, err := json.Marshal(final)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile("_data/articles.json", output, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
