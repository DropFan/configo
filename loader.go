// Package configloader : load config struct from toml, json or yaml file.
package configo

// By DropFan <DropFan@Gmail.com> @ 2017/09
// version : 0.1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/BurntSushi/toml"
)

const (
	FILE_TYPE_NONE = iota
	FILE_TYPE_TOML
	FILE_TYPE_JSON
	FILE_TYPE_YAML
)

// ConfigType config file type
type ConfigType int

// Loader config loader struct
type Loader struct {
	configFile string
	configType ConfigType
	confPtr    interface{} // &conf
}

// GetFile return config file path
func (c *Loader) GetFile() string {
	return c.configFile
}

// LoadFromFile load config from Config.FilePath
func (c *Loader) LoadFromFile(f string) error {
	var err error
	c.SetFile(f)

	switch c.configType {
	case FILE_TYPE_NONE:
		fallthrough
	case FILE_TYPE_TOML:
		err = c.LoadFromTomlFile(f)
	case FILE_TYPE_JSON:
		err = c.LoadFromJSONFile(f)
	case FILE_TYPE_YAML:
		err = c.LoadFromYamlFile(f)
	default:
		err = errors.New("unsupported config type")
	}

	return err
}

// LoadFromTomlFile load config from toml file
func (c *Loader) LoadFromTomlFile(f string) error {
	_, err := toml.DecodeFile(f, c.confPtr)
	return err
}

// LoadFromJSONFile load config from json file
func (c *Loader) LoadFromJSONFile(f string) error {
	var (
		bytes []byte
		err   error
	)

	bytes, err = ioutil.ReadFile(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, c.confPtr)

	return err
}

// LoadFromYamlFile load config from yaml file
func (c *Loader) LoadFromYamlFile(f string) error {
	var (
		bytes []byte
		err   error
	)

	bytes, err = ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, c.confPtr)
	return err
}

// LoadFromToml load config from toml string
func (c *Loader) LoadFromToml(s string) error {
	_, err := toml.Decode(s, c.confPtr)

	return err
}

// Reload reload config from config file
func (c *Loader) Reload() error {
	return c.LoadFromFile(c.configFile)
}

// SetFileType set config file type
func (c *Loader) SetFileType(t ConfigType) {
	c.configType = t
}

// SetFile set config filepath and type
func (c *Loader) SetFile(f string) error {
	var cType ConfigType

	if _, err := os.Stat(f); err != nil {
		return err
	}

	c.configFile = f

	switch strings.ToLower((path.Ext(c.configFile))) {
	case ".json":
		cType = FILE_TYPE_JSON
	case ".conf":
		fallthrough
	case ".cfg":
		fallthrough
	case ".toml":
		cType = FILE_TYPE_TOML
	case ".yaml":
		cType = FILE_TYPE_YAML
	default:
		cType = FILE_TYPE_NONE
	}
	c.configType = cType

	return nil
}

// end type Config

// New : Get an instance of config loader. `v` must be a pointer to config struct
func New(v interface{}) (*Loader, error) {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("New config loader of non-pointer %s", reflect.TypeOf(v))
	}
	if rv.IsNil() {
		return nil, fmt.Errorf("New config loader of nil %s", reflect.TypeOf(v))
	}

	c := &Loader{
		configFile: "",
		configType: FILE_TYPE_NONE,
		confPtr:    v,
	}

	return c, nil
}

// NewFromFile Get an instance of config from file. `conf` is a pointer to conf struct
func NewFromFile(file string, v interface{}) (*Loader, error) {

	c, err := New(v)

	if err != nil {
		return nil, err
	}

	c.SetFile(file)
	err = c.LoadFromFile(file)

	return c, err
}
