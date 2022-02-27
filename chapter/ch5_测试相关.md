## 内置单元测试框架

- Fail, Error: 该测试失败，该测试继续，其他测试继续执⾏
- FailNow, Fatal: 该测试失败，该测试中⽌，其他测试继续执⾏

## Benchmark

> go test -bench=. -benchmem

## BDD(Behavior Driven Development)

项⽬⽹站

https://github.com/smartystreets/goconvey


安装

go get -u github.com/smartystreets/goconvey/convey

启动 WEB UI

$GOPATH/bin/goconvey