## easyjson

[/src/test/ch8/json](../src/test/chapter/ch8/01_json)

EasyJSON 采⽤代码⽣成⽽⾮反射

安装
 go get -u github.com/mailru/easyjson/...

使⽤
 easyjson -all <结构定义>.go

## http

[/src/main/http](../src/main/chapter/ch8_02_http)

### 路由规则

- URL 分为两种，末尾是 /：表示⼀个⼦树，后⾯可以跟其他⼦路径； 末尾不是 /，表示⼀个叶⼦，固定的路径
- 以/ 结尾的 URL 可以匹配它的任何⼦路径，⽐如 /images 会匹配 /images/cute-cat.jpg
- 它采⽤最⻓匹配原则，如果有多个匹配，⼀定采⽤匹配路径最⻓的那个进⾏处理
- 如果没有找到任何匹配项，会返回 404 错误

