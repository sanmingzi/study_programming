# map

map 和 slice 一样，都是引用类型。

使用字面量创建

```go
ma := map[string]int{"a": 1, "b": 2}
```

使用内置的 make 函数创建

```go
make(map[string]int)
```

range 用来遍历 map

delete(mapName, key) 用来删除某个 k-v

```go
map_count := make(map[string]int) # initialize
count, ok := map_count[key]
if ok {
  count += 1
} else {
  count = 1
}
```
