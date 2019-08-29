[![Build Status](https://travis-ci.org/raelga/yodo.svg?branch=master)](https://travis-ci.org/raelga/yodo)
[![Go Report Card](https://goreportcard.com/badge/github.com/raelga/yodo)](https://goreportcard.com/report/github.com/raelga/yodo)
[![Coverage Status](https://coveralls.io/repos/github/raelga/yodo/badge.svg?branch=master)](https://coveralls.io/github/raelga/yodo?branch=master)
[![GitHub](https://img.shields.io/github/license/raelga/yodo)](https://github.com/raelga/yodo/blob/master/LICENSE)

# yodo

Simple GoLang TODO application

This application is just a playground to learn:

- TDD with GoLang
- github.com/spf13/cobra
- github.com/spf13/viper
- gopkg.in/yaml.v2
- GitHub Actions


```
$ go build 
```

```yaml
$ ./yodo get
default:
Empty list
```

```yaml
$ ./yodo add Add verbose mode
add task "Add verbose mode" (false) to file$ 
```

```yaml
$ ./yodo get
default:
- [0] "Add verbose mode" (false)
$ ./yodo do 0
Task 0 done.
```

```yaml
$ ./yodo get
default:
- [0] "Add verbose mode" (true)
```

```yaml
$ cat ~/.yodo/default.yaml 
id: 0
name: default
tasks:
- id: 0
  status: true
  description: Add verbose mode
```
