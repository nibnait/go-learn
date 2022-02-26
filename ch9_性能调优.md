## 性能调优

[/src/main/tools](./src/main/ch8_02_http)

### 准备⼯作

- 安装 graphviz

> brew install graphviz

- 将 $GOPATH/bin 加⼊ $PATH

> Mac OS: 在 .bash_profile 中修改路径

- 安装 go-torch

> go get -u github.com/uber/go-torch
>
> 下载并复制 flamegraph.pl （https://github.com/brendangregg/FlameGraph）⾄ $GOPATH/bin 路径下 
>
> 将 $GOPATH/bin 加⼊ $PATH

### 通过⽂件⽅式输出 Proﬁle

- 灵活性⾼，适⽤于特定代码段的分析
- 通过⼿动调⽤ runtime/pprof 的 API
- API 相关⽂档 https://studygolang.com/static/pkgdoc/pkg/runtime_pprof.htm
- go tool pprof [binary] [binary.prof]

> go build prof.go
>
> ls
>
> ./prof
>
> ls
>
> go tool pprof prof cpu.prof
>
> > top
> >
> > list fillMatrix
> >
> > svg
> >
> > exit
>
> go-torch cpu.prof（查看火炬图）
>
> go tool pprof prof mem.prof
>
> // 先gc 在 dump，，再重新观察
>
> go tool pprof prof cpu.prof
>
> > list



### 通过 http 方式输出 Profile

- 简单，适合于持续性运⾏的应⽤
- 在应⽤程序中导⼊ import _ "net/http/pprof"，并启动 http server 即可
- http://<host>:<port>/debug/pprof/
- go tool pprof http://<host>:<port>/debug/pprof/profile?seconds=10 （默认值为30秒）
- go-torch -seconds 10 http://<host>:<port>/debug/pprof/proﬁle

> go run fb_server.go
>
> http://127.0.0.1:8081/fb



> go build fb_server.go
>
> ./fb_server
>
> go tool pprof prof cpu.prof
>
> > top
> >
> > list fillMatrix
> >
> > svg
>
> go tool pprof prof cpu.prof
>
> > list

## 性能调优示例

### 常⻅分析指标

- Wall Time
- CPU Time
- Block Time
- Memory allocation
- GC times/time spent

> cd src/test/ch9
>
> go test -bench=.
>
> go test -bench=. -cpuprofile=cpu.prof
>
> go tool pprof cpu.prof
>
> > top -cum
> >
> > top
> >
> > list processRequest
>
> go test -bench=. -memprofile=mem.prof
>
> go tool pprof mem.prof

> go help testflag

### 别让性能被“锁”住

- 减少锁的影响范围
- 减少发⽣锁冲突的概率
- sync.Map 
  - 适合读多写少，且 Key 相对稳定的环 境
  - 采⽤了空间换时间的⽅案，并且采⽤指针的⽅式间接实现值的映射，所以存储空间会较 built-in map ⼤
- **ConcurrentMap** 
  - 适⽤于读写都很频繁的情况
- 避免锁的使⽤
- LAMX **Disruptor**：https://martinfowler.com/articles/lmax.html

### GC 友好的代码

#### 避免内存分配和复制

- 复杂对象尽量传递引⽤
  - 数组的传递
  - 结构体传递
- 初始化⾄合适的⼤⼩
  - ⾃动扩容是有代价的
- 复⽤内存

#### 打开 GC ⽇志

只要在程序执⾏之前加上环境变量 GODEBUG=gctrace=1，

> GODEBUG=gctrace=1 go test -bench=.
>
> GODEBUG=gctrace=1 go run main.go

⽇志详细信息参考： https://godoc.org/runtime

> cd 03_gc_friendly
>
> go test -bench=BenchmarkPassingArrayWithRef -trace=trace_ref.out
>
> go test -bench=BenchmarkPassingArrayWithValue -trace=trace_val.out
>
> ls
>
> go tool trace trace_ref.out
>
> go tool trace trace_val.out

## ⾼效的字符串连接

strings.Builder

## 面向错误的设计

### 隔离错误 - 设计

微内核

### 隔离错误 - 部署

微服务

### 重用 vs 隔离

逻辑结构的重用 + 部署结构的隔离

### 冗余

主从

### 限流

### 慢响应

A quick rejection is better than a slow response.

给阻塞操作都加上⼀个期限（配置接口、服务 级别的超时时间）

### 错误传递

断路器（服务降级）

## 面向恢复的设计

### 健康检查

- 注意僵⼫进程
- 池化资源耗尽
- 死锁

### 构建可恢复的系统

- 拒绝单体系统
- ⾯向错误和恢复的设计
- 在依赖服务不可⽤时，可以继续存活 快速启动
- ⽆状态

### 与客户端协商

相同请求，一段时间内 不要发送过来

# 相关开源项⽬

[https://github.com/Netﬂix/chaosmonkey](https://github.com/Netﬂix/chaosmonkey)

[https://github.com/easierway/service_decorators/blob/master/README.md](https://github.com/easierway/service_decorators/blob/master/README.md)
