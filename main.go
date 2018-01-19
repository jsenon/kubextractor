// Package main Kubeextractor.
//
// the purpose of this application is to provide an Kube Configuration Extraction
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
}

func main() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		fmt.Print("Error:", err)
	}
	fmt.Println(conf)
}
