你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
