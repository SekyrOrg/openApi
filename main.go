package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type OpenAPISpec struct {
	OpenAPI    string                 `yaml:"openapi"`
	Info       map[string]interface{} `yaml:"info"`
	Paths      map[string]interface{} `yaml:"paths"`
	Components map[string]interface{} `yaml:"components"`
}

func main() {
	sourceFile := flag.String("s", "specs/creator.yaml", "OpenApi Spec file to merge from (source)")
	targetFile := flag.String("t", "specs/server.yaml", "OpenApi Spec file to merge into (target)")
	infoFile := flag.String("i", "specs/gateway_info.yaml", "OpenApi file with info section to use")
	outputFile := flag.String("o", "specs/gateway.yaml", "Where to save the merged OpenApi Spec file")
	ignoreKeys := flag.String("ignore", "/healthz", "comma separated list of keys to ignore when merging")
	flag.Parse()
	// convert ignoreKeys to map
	ignoreKeysMap := make(map[string]any)
	for _, key := range strings.Split(*ignoreKeys, ",") {
		ignoreKeysMap[key] = nil
	}
	if err := Run(*sourceFile, *targetFile, *infoFile, *outputFile, ignoreKeysMap); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(sourceFile, targetFile, infoFile, resultFile string, ignoreKeys map[string]any) error {
	sourceData, targetData, err := ParseSpecFiles(sourceFile, targetFile)
	if err != nil {
		return fmt.Errorf("error parsing spec files: %s", err)
	}
	// mergeOpenApiSpec the components recursively
	mergeOpenApiSpec(sourceData, targetData, ignoreKeys)

	if infoFile != "" {
		infoData, err := ParseSpecFile(infoFile)
		if err != nil {
			return fmt.Errorf("error parsing info file: %s", err)
		}
		// overwrite the info section
		targetData.Info = infoData.Info
	}

	// write the result to a file
	if err := SaveResult(resultFile, targetData); err != nil {
		return fmt.Errorf("error saving result: %s", err)
	}
	return nil
}

func SaveResult(resultFile string, targetData *OpenAPISpec) error {
	file, err := os.Create(resultFile)
	if err != nil {
		return fmt.Errorf("error creating result file: %s", err)
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	defer encoder.Close()
	if err = encoder.Encode(targetData); err != nil {
		return fmt.Errorf("error writing result file: %s", err)
	}
	return nil
}

func ParseSpecFile(specFile string) (*OpenAPISpec, error) {
	file, err := os.Open(specFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	data := &OpenAPISpec{}
	if err = decoder.Decode(data); err != nil {
		return nil, fmt.Errorf("error decoding file: %s", err)
	}
	return data, nil
}

func ParseSpecFiles(sourceFile string, targetFile string) (*OpenAPISpec, *OpenAPISpec, error) {
	sourceData, err := ParseSpecFile(sourceFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing source file: %s", err)
	}
	targetData, err := ParseSpecFile(targetFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing target file: %s", err)
	}
	return sourceData, targetData, nil
}

func mergeOpenApiSpec(source *OpenAPISpec, target *OpenAPISpec, ignoreKeys map[string]any) {
	merge(source.Components, target.Components, ignoreKeys)
	merge(source.Paths, target.Paths, ignoreKeys)
}

func merge(source map[string]interface{}, target map[string]interface{}, ignoreKeys map[string]any) {
	for key, value := range source {
		if _, ok := target[key]; ok {
			// if the key exists in the target, check if it's a map
			if _, ok := value.(map[string]interface{}); !ok {
				// if it's not a map, just overwrite the value
				target[key] = value
				continue
			}
			// mergeOpenApiSpec the components recursively
			merge(value.(map[string]interface{}), target[key].(map[string]interface{}), ignoreKeys)
		} else {
			// check if the key is in the ignore list
			if _, exists := ignoreKeys[key]; exists {
				continue
			}
			target[key] = value
		}
	}
}
