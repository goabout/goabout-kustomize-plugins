package main

import (
	"go.mozilla.org/sops/decrypt"
	"sigs.k8s.io/kustomize/k8sdeps/kv"
	"sigs.k8s.io/kustomize/pkg/fs"
	"sigs.k8s.io/kustomize/pkg/loader"
)

type sopsDotEnv struct{}

var KVSource sopsDotEnv

func (p *sopsDotEnv) Get(root string, args []string) (map[string]string, error) {
	ldr, err := loader.NewLoader(root, fs.MakeRealFS())
	if err != nil {
		return nil, err
	}

	allKvs := make(map[string]string)
	for _, fPath := range args {
		content, err := ldr.Load(fPath)
		if err != nil {
			return nil, err
		}

		data, err := decrypt.Data(content, "dotenv")
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
