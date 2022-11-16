package config

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const configFile = "./testdata/config_test.yml"

func TestInit(t *testing.T) {
	Init(configFile)
}

func TestGetString(t *testing.T) {
	Init(configFile)
	appMode := GetString("App.Mode")
	assert.Equal(t, "Debug", appMode)
}

func TestGetInt(t *testing.T) {
	Init(configFile)
	serverPort := GetInt("TemporalServer.Port")
	assert.Equal(t, 7233, serverPort)
}

func ReadYamlFile(t *testing.T, fileName string, data interface{}) {
	// 读取 yaml 到结构体
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal(file, data)
	if err != nil {
		t.Fatal(err)
	}
}

func WriteYamlFile(t *testing.T, fileName string, data interface{}) {
	out, err := yaml.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(fileName, out, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
