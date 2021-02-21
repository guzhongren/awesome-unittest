## 覆盖率报告

```shell
go test -coverprofile=coverage.out
```

## 查看报告

```shell
go tool cover -html=coverage.out
```