package application

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
)

func getTempPath(directory string) string {
	absPath, err := filepath.Abs(directory)
	if err != nil {
		panic(err)
	}
	return absPath
}

var (
	Config      = loadConfig()
	configPath  = "config.json"
	defaultConf = Cgf{
		Port:      3000,
		Debug:     true,
		MongoDB:   "gamechanger",
		MongoPort: 27017,
		MongoHost: "127.0.0.1",
	}
)

// Cgf set required tag form member to force it in config
type (
	Cgf struct {
		Port      int    `json:"port"`
		MongoDB   string `json:"mongo_db" required:"true"`
		MongoHost string `json:"mongo_host" required:"true"`
		MongoPort int    `json:"mongo_port" required:"true"`
		Debug     bool   `json:"debug"`
	}
)

func (c *Cgf) requiredColumnExists() bool {
	allExists := true
	cType := reflect.TypeOf(c)
	cVal := reflect.ValueOf(c)
	if cVal.Kind() == reflect.Ptr {
		cVal = cVal.Elem()
		cType = cType.Elem()
	}
	for i := 0; i < cVal.NumField(); i++ {
		field := cType.Field(i)
		val, ok := field.Tag.Lookup("required")
		if !ok || val == "false" {
			continue
		}
		value := cVal.Field(i)
		switch value.Type().Kind() {
		case reflect.Bool:
			continue
		case reflect.String:
			if value.String() == "" {
				allExists = false
			}
		case reflect.Int:
			if value.Int() == 0 {
				allExists = false
			}
		default:
			allExists = false
		}
	}
	return allExists
}

func loadConfig() *Cgf {
	cgf := &defaultConf
	file, err := ioutil.ReadFile(configPath)
	if err == nil {
		json.Unmarshal(file, cgf)
	}
	if !cgf.requiredColumnExists() {
		log.Fatal("Loading Config required faild")
	}
	return cgf
}
