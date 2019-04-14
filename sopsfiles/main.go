package main

import (
	"fmt"
	"go.mozilla.org/sops/cmd/sops/common"
	"go.mozilla.org/sops/decrypt"
	"sigs.k8s.io/kustomize/k8sdeps/kv"
)

type sopsFiles struct{}

var KVSource sopsFiles

func (p *sopsFiles) Get(root string, args []string) (map[string]string, error) {
	kvs := make(map[string]string)
	for _, s := range args {
		k, fPath, err := kv.ParseFileSource(s)
		if err != nil {
			return nil, err
		}
		format := formatForPath(fPath)
		data, err := decrypt.File(fmt.Sprintf("%s/%s", root, fPath), format)
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
