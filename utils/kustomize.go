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

	log "github.com/sirupsen/logrus"
	"encoding/json"

	"sigs.k8s.io/kustomize/kustomize/v4/commands/build"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

type T struct {
	err  error
	path string
}

type Target struct {
	Kind      string
	Name      string
	Namespace string
}

type Patch struct {
	Patch  string
	Target Target
}
type Patches []Patch

func checkPatch(patches []byte) (bool, error) {
	if err := json.Unmarshal(patches, &Patches{}); err != nil {
		return false, err
	}
	return true, nil
}

func Kustomize(path string) (string, error) {
	fs := filesys.MakeFsOnDisk()

	buffy := new(bytes.Buffer)
	cmd := build.NewCmdBuild(fs, build.MakeHelp("foo", "bar"), buffy)
	if err := cmd.RunE(cmd, []string{path}); err != nil {
		return "", err
	}

	return buffy.String(), nil
}

func RunKustomizations(bs []B, o bool) {
	ts := []T{}

	for _, b := range bs {
		path := b.Path
		y := getYamlObj(path)
		updateKustomization(b, y)
		output, err := Kustomize(path)
		if err != nil {
			ts = append(ts, T{
				err:  err,
				path: path,
			})
		} else if o {
			log.Print(output)
		}
		removePatches(path, y)
	}

	if len(ts) == 0 {
		log.Print("Successfully reconciled all kustomizations")
	}
	for _, t := range ts {
		log.Errorf("Flux failed to reconcile %s: %v", t.path, t.err)
	}
}
