package main

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"gopkg.in/yaml.v3"
	"os"
)

type OpenAPISpec struct {
	OpenAPI    string                   `yaml:"openapi"`
	Info       map[string]interface{}   `yaml:"info"`
	Paths      map[string]interface{}   `yaml:"paths"`
	Components map[string]interface{}   `yaml:"components"`
	Tags       []map[string]interface{} `yaml:"tags"`
	Servers    []map[string]interface{} `yaml:"servers"`
}

type CliArgs struct {
	SourceFile  string
	TargetFiles goflags.StringSlice
	IgnoredKeys goflags.StringSlice
	Merge       goflags.StringSlice
	OutputFile  string
}

func main() {
	args := CliArgs{}
	flagSet := goflags.NewFlagSet()
	flagSet.StringVarP(&args.SourceFile, "source", "s", "specs/gateway.yaml", "OpenApi Spec file to merge from (source)")
	flagSet.StringSliceVarP(&args.TargetFiles, "target", "t", []string{"specs/server.yaml", "specs/creator.yaml"}, "OpenApi Spec files to merge into (target)", goflags.CommaSeparatedStringSliceOptions)
	flagSet.StringSliceVarP(&args.IgnoredKeys, "ignore", "i", []string{"/healthz"}, "Keys to ignore when merging", goflags.CommaSeparatedStringSliceOptions)
	flagSet.StringVarP(&args.OutputFile, "output", "o", "specs/gateway_merged.yaml", "Output file for the merged OpenApi Spec")
	flagSet.StringSliceVarP(&args.Merge, "merge", "m", []string{"paths", "components", "tags", "servers"}, "Keys to merge when merging. Keys are separated by a comma. If no keys are specified, all keys will be merged. Available keys are 'paths', 'components', 'tags', 'servers'", goflags.CommaSeparatedStringSliceOptions)
	flagSet.Parse()
	// convert ignoreKeys to map
	ignoreKeys := make(map[string]any)
	for _, key := range args.IgnoredKeys {
		ignoreKeys[key] = nil
	}
	// convert merge to map
	mergeKeys := make(map[string]any)
	for _, key := range args.Merge {
		mergeKeys[key] = nil
	}

	if err := Run(args.SourceFile, args.TargetFiles, args.OutputFile, ignoreKeys, mergeKeys); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run(sourceFile string, targetFiles []string, outFile string, ignoreKeys map[string]any, mergeKeys map[string]any) error {
	sourceData, targetData, err := ParseSpecFiles(sourceFile, targetFiles)
	if err != nil {
		return fmt.Errorf("error parsing spec files: %s", err)
	}
	// mergeOpenApiSpec the components recursively
	for _, data := range targetData {
		mergeOpenApiSpec(sourceData, data, ignoreKeys, mergeKeys)
	}

	// write the result to a file
	if err := SaveResult(outFile, sourceData); err != nil {
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

func ParseSpecFiles(sourceFile string, targetFiles []string) (*OpenAPISpec, []*OpenAPISpec, error) {
	sourceData, err := ParseSpecFile(sourceFile)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing source file:%s, error: %s", sourceFile, err)
	}
	var targetList []*OpenAPISpec
	for _, targetFile := range targetFiles {
		targetData, err := ParseSpecFile(targetFile)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing target file:%s, error: %s", targetFile, err)
		}
		targetList = append(targetList, targetData)
	}
	return sourceData, targetList, nil
}

func mergeOpenApiSpec(source, target *OpenAPISpec, ignoreKeys, mergeKeys map[string]any) {
	if _, ok := mergeKeys["servers"]; ok {
		mergeList(target.Servers, source.Servers, ignoreKeys)
	}
	if _, ok := mergeKeys["components"]; ok {
		merge(target.Components, source.Components, ignoreKeys)
	}
	if _, ok := mergeKeys["paths"]; ok {
		merge(target.Paths, source.Paths, ignoreKeys)
	}
	if _, ok := mergeKeys["tags"]; ok {
		mergeList(target.Tags, source.Tags, ignoreKeys)
	}
}
func mergeList(source []map[string]interface{}, target []map[string]interface{}, ignoreKeys map[string]any) {
	for _, value := range source {
		target = append(target, value)
	}
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
