# Custom Linter plagin for golangci-lint

<h2>How to build linter itself</h2>

```
go build cmd/main.go
```

<h4>How to test it</h4>

```
go test ./analyzer/rules_test.go

go test analazer_test.go

./main ./testdata/src/foo/example.go

```

<h2>How to build plugin<h2>

```
go build -buildmode=plugin -o ./plugin/customlinter.so ./plugin/main.go 


```
