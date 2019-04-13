package main

import (
	"fmt"
	"go.mozilla.org/sops/cmd/sops/common"
	"go.mozilla.org/sops/decrypt"
	"strings"
)

type sopsFiles struct{}

var KVSource sopsFiles

func (p *sopsFiles) Get(root string, args []string) (map[string]string, error) {
	allKvs := make(map[string]string)

	for _, arg := range args {
		split := strings.SplitN(arg, "=", 2)
		if len(split) < 2 {
			return nil, fmt.Errorf("should be NAME=FILENAME: %s", arg)
		}
		format := formatForPath(split[1])
		data, err := decrypt.File(fmt.Sprintf("%s/%s", root, split[1]), format)
		if err != nil {
			return nil, err
		}

		allKvs[split[0]] = string(data)
	}

	return allKvs, nil
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
