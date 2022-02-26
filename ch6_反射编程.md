 - reﬂect.TypeOf 返回类型 (reﬂect.Type)
 - reﬂect.ValueOf 返回值 (reﬂect.Value)
 - 可以从 reﬂect.Value 获得类型
 - 通过 kind 的来判断类型

## 利⽤反射编写灵活的代码

按名字访问结构的成员

```go
reflect.ValueOf(*e).FieldByName("Name")
```

按名字访问结构的⽅法

```go
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
```