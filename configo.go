// Package configo : load config struct from toml, json or yaml file.
// By DropFan <DropFan@Gmail.com> @ 2017/09
// version 0.1
package configo

import (
	"fmt"
	"reflect"
)

// New : Get an instance of config loader. `configPtr` must be a pointer to config struct
func New(configPtr interface{}) (*Loader, error) {

	rv := reflect.ValueOf(configPtr)
	if rv.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("New config loader of non-pointer %s", reflect.TypeOf(configPtr))
	}
	if rv.IsNil() {
		return nil, fmt.Errorf("New config loader of nil %s", reflect.TypeOf(configPtr))
	}

	c := &Loader{
		configFile: "",
		configType: FileTypeNone,
		confPtr:    configPtr,
	}

	return c, nil
}

// NewFromFile Get an instance of config loader and load config from file. `configPtr` is a pointer to config struct
func NewFromFile(file string, configPtr interface{}) (*Loader, error) {

	c, err := New(configPtr)

	if err != nil {
		return nil, err
	}

	c.SetFile(file)
	err = c.LoadFromFile(file)

	return c, err
}
