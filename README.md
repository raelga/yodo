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