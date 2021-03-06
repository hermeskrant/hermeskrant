package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
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

var timeString = "2006-01-02"

const dataFolder = "_data"
const articlesfolder = dataFolder + "/articles"
const articlespagedata = dataFolder + "/articles.json"

func main() {
	infos, err := ioutil.ReadDir(articlesfolder)
	if err != nil {
		fmt.Println(err.Error())
	}

	var files = make([]interface{}, len(infos))
	var final = make(map[string]interface{}, 1)

	for i, info := range infos {
		file, err := ioutil.ReadFile(filepath.Join(articlesfolder, info.Name()))
		if err != nil {
			fmt.Println(err.Error())
		}

		err = json.Unmarshal(file, &files[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		files[i].(map[string]interface{})["path"] = strings.Replace(info.Name(), ".json", "", -1)

		t := strings.Split(files[i].(map[string]interface{})["path"].(string), "-")
		year, err := strconv.Atoi(t[0])
		if err != nil {
			fmt.Println(err.Error())
		}

		month, err := strconv.Atoi(t[1])
		if err != nil {
			fmt.Println(err.Error())
		}

		day, err := strconv.Atoi(t[2])
		if err != nil {
			fmt.Println(err.Error())
		}

		files[i].(map[string]interface{})["date"] = fmt.Sprintf("%v %v %v", day, months[month-1], year)

		outputJSON, err := json.Marshal(files[i])
		if err != nil {
			fmt.Println(err.Error())
		}

		err = ioutil.WriteFile(articlesfolder+info.Name(), outputJSON, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	final["Articles"] = reverse(files)

	output, err := json.Marshal(final)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile(articlespagedata, output, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func reverse(numbers []interface{}) []interface{} {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
