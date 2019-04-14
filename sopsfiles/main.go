package main

import (
	"go.mozilla.org/sops/cmd/sops/common"
	"go.mozilla.org/sops/decrypt"
	"sigs.k8s.io/kustomize/k8sdeps/kv"
	"sigs.k8s.io/kustomize/pkg/fs"
	"sigs.k8s.io/kustomize/pkg/loader"
)

type sopsFiles struct{}

var KVSource sopsFiles

func (p *sopsFiles) Get(root string, args []string) (map[string]string, error) {
	ldr, err := loader.NewLoader(root, fs.MakeRealFS())
	if err != nil {
		return nil, err
	}

	kvs := make(map[string]string)
	for _, s := range args {
		k, fPath, err := kv.ParseFileSource(s)
		if err != nil {
			return nil, err
		}

		content, err := ldr.Load(fPath)
		if err != nil {
			return nil, err
		}

		format := formatForPath(fPath)
		data, err := decrypt.Data(content, format)
		if err != nil {
			return nil, err
		}

		kvs[k] = string(data)
	}
	return kvs, nil
}

func formatForPath(path string) string {
	if common.IsYAMLFile(path) {
		return "yaml"
	} else if common.IsJSONFile(path) {
		return "json"
	} else if common.IsEnvFile(path) {
		return "dotenv"
	}
	return "binary"
}
