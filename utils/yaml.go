/*
Copyright © 2021 Josh Hill

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type B struct {
	Path    string
	Patches interface{}
}

func SplitYaml(doc string) []B {
	b := []byte(doc)
	dec := yaml.NewDecoder(bytes.NewReader(b))

	var res []B
	for {
		var value map[string]interface{}

		err := dec.Decode(&value)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		spec, ok := value["spec"].(map[interface{}]interface{})
		if !ok {
			log.Fatal(err)
		}

		b := B{
			Path: spec["path"].(string),
			Patches: spec["patches"],
		}

		res = append(res, b)
	}
	return res
}

func getYamlObj(path string) map[string]interface{} {
	log.Debug(fmt.Sprintf("opening kustomization.yaml at %s", path))
	f, err := ioutil.ReadFile(path + "/kustomization.yaml")
	if err != nil {
		log.Fatalf("getYamlObj: %v", err)
	}
	var y map[string]interface{}
	err = yaml.Unmarshal(f, &y)
	if err != nil {
		log.Fatal(err)
	}

	return y
}

func updateKustomization(b B, y map[string]interface{}) {
	y["patches"] = b.Patches
	bytes, err := yaml.Marshal(y)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(b.Path + "/kustomization.yaml", bytes, 0644); err != nil {
		log.Fatal(err)
	}
}

func removePatches(path string, y map[string]interface{}) {
	delete(y, "patches")
	bytes, err := yaml.Marshal(y)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(path + "/kustomization.yaml", bytes, 0644); err != nil {
		log.Fatal(err)
	}
}
