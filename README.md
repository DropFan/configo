# Configo

[![Build Status](https://travis-ci.org/DropFan/configo.svg?branch=master)](https://travis-ci.org/DropFan/configo)[![Go Report Card](https://goreportcard.com/badge/github.com/DropFan/configo)](https://goreportcard.com/report/github.com/DropFan/configo)

A lightweight package to loading config file for golang. (support toml, json, yaml.)

## Installation

`go get -u github.com/DropFan/configo`

## Usage
Click [here](https://github.com/DropFan/configo/tree/master/examples) to get examples.
```go
import (
    "github.com/DropFan/configo"
)
var (
    demoConf     *DemoConfigStruct
    configLoader *configo.Loader
)
configLoader, err = configo.NewFromFile(f, demoConf)
```

## Contacts

Author: Tiger(DropFan)

Email: <DropFan@Gmail.com>

Wechat: DropFan

Telegram: [DropFan](https://telegram.me/DropFan)

[https://about.me/DropFan](https://about.me/DropFan)

## License

[MIT](https://github.com/DropFan/configo/blob/master/LICENSE)
