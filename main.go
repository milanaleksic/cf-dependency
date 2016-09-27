package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

type Resource struct {
	DependsOn interface{} `json:"DependsOn,omitempty"`
}

type Template struct {
	Resources map[string]Resource `json:"Resources"`
}

func parse(file string) (parsedTemplate Template) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &parsedTemplate)
	if err != nil {
		panic(err)
	}
	return parsedTemplate
}

func main() {
	fmt.Println("digraph dependencies {")
	var knownTemplates = os.Args
	for i, template := range knownTemplates {
		if i == 0 {
			continue
		}
		parsedTemplate := parse(template)
		for resourceName, resource := range parsedTemplate.Resources {
			switch t := resource.DependsOn.(type) {
			case string:
				fmt.Printf("%s -> %s;\n", t, resourceName)
			case []interface {}:
				for _, singleDependency := range t {
					fmt.Printf("%s -> %s;\n", singleDependency, resourceName)
				}
			default:
				fmt.Printf("%s;\n", resourceName)
			}
		}
	}
	fmt.Println("}")
}
