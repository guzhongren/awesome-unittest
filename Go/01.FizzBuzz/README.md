# TDD

## 测试提示

### 执行当前目录下的测试

```shell
$ go test .
...
```

### 生成测试报告后用浏览器打开

```shell
$ go test -coverprofile=c.out
...
$ go tool cover -html=c.out
...
```

### 性能测试

```shell
$ go test -bench .
...
```