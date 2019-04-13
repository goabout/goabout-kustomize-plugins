package main

import (
	"fmt"
	"go.mozilla.org/sops/decrypt"
	"sigs.k8s.io/kustomize/k8sdeps/kv"
)

type sopsDotEnv struct{}

var KVSource sopsDotEnv

func (p *sopsDotEnv) Get(root string, args []string) (map[string]string, error) {
	allKvs := make(map[string]string)
	for _, arg := range args {
		data, err := decrypt.File(fmt.Sprintf("%s/%s", root, arg), "dotenv")
		if err != nil {
			return nil, err
		}

		kvs, err := kv.KeyValuesFromLines(data)
		if err != nil {
			return nil, err
		}

		for _, pair := range kvs {
			allKvs[pair.Key] = pair.Value
		}
	}
	return allKvs, nil
}
