package utils

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//YamlFileToMap Convert yaml fil to map
func YamlFileToMap(configFile string) (*map[interface{}]interface{}, error) {
	yamlMap := make(map[interface{}]interface{})
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, yamlMap)
	if err != nil {
		return nil, err
	}
	return &yamlMap, nil
}

//MapToYamlString convert map to yaml
func MapToYamlString(yamlMap map[string]interface{}) (string, error) {
	lock.Lock()
	defer lock.Unlock()
	d, err := yaml.Marshal(&yamlMap)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func ParseJsonFile(filepath string) ([]byte, error) {
	jsonFile, err := os.Open(filepath)
	defer func() { err = jsonFile.Close() }()
	if err != nil {
		return nil, err
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

//struct 转 map
func Struct2Map(content interface{}) (map[string]interface{}, error) {
	var structMap map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		return nil, err
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&structMap); err != nil {
			return nil, err
		} else {
			for k, v := range structMap {
				structMap[k] = v
			}
		}
	}
	return structMap, nil
}
