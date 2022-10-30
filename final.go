package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	//"strings"
)

func main() {

			// Reading the mapping.json file and decoding
	jsonFile, _ := os.Open("mapping.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result_map map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result_map)
			// Dynamically decoding every objects of mapping.json
	mapping := map[string]map[string]interface{}{}
	for k,v := range result_map{
		switch c := v.(type){
		case map[string]interface{}:
			mapping[k] = v.(map[string]interface{})
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}

			//Reading the data.json file and decoding
	empJson, _ := os.Open("data.json")
	defer empJson.Close()
	byteValue_2, _ := ioutil.ReadAll(empJson)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue_2), &result)
			// Dynamically decoding every object of data.json
	data := map[string]map[string]interface{}{}
	for k, v := range result {
		switch c := v.(type) {
		case map[string]interface{}:
			data[k] = v.(map[string]interface{})
		case []interface{}:
			for _, item := range v.([]interface{}) {
				data[k] = item.(map[string]interface{})
			}
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}

			// Taking another map to store keys and values
	CopiedMap:= make(map[string]interface{})
	for v := range result_map{
		for item := range mapping[v] {

			value := (strings.Split(fmt.Sprint((mapping[v][item])),"."))
			CopiedMap[item] = data[value[0]][value[1]]
			fmt.Printf("%v : %v\n", item, CopiedMap[item])
			fmt.Println(" ")
			
		}
	}

}
